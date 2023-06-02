import { CheckIcon, EditIcon } from '@chakra-ui/icons';
import {
  Box,
  Flex,
  Grid,
  GridItem,
  Heading,
  Highlight,
  HStack,
  IconButton,
} from '@chakra-ui/react';
import EmojiPicker from '@components/inputs/emoji/EmojiPicker';
import MultiSelect from '@components/inputs/multi-select/MultiSelect';
import Select from '@components/inputs/select/Select';
import TextInput from '@components/inputs/text/TextInput';
import { useNoteStore } from '@lib/stores';
import { NoteStatus, Tag } from '@lib/types';
import cn from 'classnames';
import { Emoji, EmojiStyle } from 'emoji-picker-react';
import { Formik } from 'formik';
import { useState } from 'react';

import s from './NoteMetadataSection.module.scss';

export interface INoteMetadataSection {
  data: {
    tags?: Tag[];
  };
  disabled?: boolean;
}

const NoteMetadataSection: React.FC<INoteMetadataSection> = ({
  data,
  disabled = false,
}) => {
  const [tags, _setTags] = useState<Tag[]>(data.tags ?? []);
  const [isEditable, setIsEditable] = useState(true);
  const [isLoading, setIsLoading] = useState(false);

  const { note, updateNote } = useNoteStore();

  if (!note) return <></>;

  const initialValues: {
    title: string;
    tags: string[];
    emoji: string;
    status: NoteStatus;
  } = {
    title: note.title,
    tags: note.tags.map((s) => s.id),
    emoji: note.emoji,
    status: note.status,
  };

  return (
    <Formik
      initialValues={initialValues}
      onSubmit={(values, _action) => {
        updateNote(values);
      }}
    >
      {(props) => (
        <form
          onSubmit={props.handleSubmit}
          className={cn(
            disabled ? s.root : '',
            `${disabled ? 'w-[80%]' : 'w-full'}`
          )}
        >
          <Box
            width={'full'}
            padding={'5'}
            borderRadius={'md'}
            display={'flex'}
            flexDir={'column'}
            alignItems={`${disabled ? 'center' : ''}`}
            backgroundColor={'foreground'}
            zIndex={10}
            position={'relative'}
          >
            <HStack marginBottom={5} display={'flex'} alignItems={'center'}>
              {!disabled ? (
                <Heading as={'h3'} size={'md'}>
                  <Highlight
                    query={'metadata'}
                    styles={{
                      px: '2',
                      py: '1',
                      rounded: 'full',
                      bg: 'red.100',
                    }}
                  >
                    Metadata
                  </Highlight>
                </Heading>
              ) : (
                <Flex alignItems={'center'} gap={5}>
                  <Emoji
                    unified={note.emoji}
                    size={35}
                    emojiStyle={EmojiStyle.TWITTER}
                  />
                  <Heading color={'white'}>{note.title}</Heading>
                </Flex>
              )}
              {!disabled && (
                <Flex className={'!ml-auto'} alignItems={'center'} gap={5}>
                  {isEditable ? (
                    <IconButton
                      aria-label={'Update Metadata'}
                      icon={<CheckIcon />}
                      isLoading={isLoading}
                      onClick={() => {
                        setIsLoading(true);
                        props.handleSubmit();
                        setIsLoading(false);
                        setIsEditable(false);
                      }}
                    />
                  ) : (
                    <IconButton
                      className={'!ml-auto'}
                      aria-label={'Edit Metadata'}
                      icon={<EditIcon />}
                      isLoading={isLoading}
                      onClick={() => setIsEditable(true)}
                    />
                  )}
                </Flex>
              )}
            </HStack>
            <Grid
              gridTemplateColumns={
                disabled
                  ? { base: '1fr', md: '1fr 1fr' }
                  : { base: '1fr', md: '3fr 2fr 1fr' }
              }
              gap={'5'}
              alignItems={'start'}
              justifyContent={'start'}
            >
              {!disabled && (
                <GridItem
                  display={'grid'}
                  gridTemplateColumns={'1fr max-content'}
                  gap={2}
                >
                  <TextInput
                    value={props.values.title}
                    onChange={(evt) =>
                      props.setFieldValue('title', evt.target.value)
                    }
                    label={'Title'}
                    placeholder={'Search problems'}
                  />
                  <EmojiPicker
                    value={props.values.emoji}
                    onChange={(emojiData, _evt) => {
                      props.setFieldValue('emoji', emojiData.unified);
                    }}
                    disabled={disabled || !isEditable}
                  />
                </GridItem>
              )}
              <GridItem>
                <MultiSelect
                  data={
                    disabled
                      ? note.tags.map((t) => ({
                          value: t.id,
                          label: t.label,
                          color: t.color,
                        }))
                      : tags.map((t) => ({
                          value: t.id,
                          label: t.label,
                          color: t.color,
                        }))
                  }
                  value={props.values.tags}
                  onChange={(value) => props.setFieldValue('tags', value)}
                  label={'Tags'}
                  placeholder={'Select tags'}
                  disabled={disabled || !isEditable}
                />
              </GridItem>
              <GridItem>
                <Select
                  data={[
                    {
                      value: 'DRAFT',
                      label: 'Draft',
                      color: 'yellow',
                    },
                    {
                      value: 'PUBLISHED',
                      label: 'Published',
                      color: 'teal',
                    },
                  ]}
                  value={props.values.status}
                  label={'Status'}
                  onChange={(value) => props.setFieldValue('status', value)}
                  placeholder={'Select status'}
                  disabled={disabled || !isEditable}
                />
              </GridItem>
            </Grid>
          </Box>
        </form>
      )}
    </Formik>
  );
};

export default NoteMetadataSection;
