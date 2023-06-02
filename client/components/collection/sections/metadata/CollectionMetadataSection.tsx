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
  Input,
  Text,
  Textarea,
} from '@chakra-ui/react';
import EmojiPicker from '@components/inputs/emoji/EmojiPicker';
import MultiSelect from '@components/inputs/multi-select/MultiSelect';
import Select from '@components/inputs/select/Select';
import TagLabel from '@components/labels/tag/TagLabel';
import { useCollectionStore } from '@lib/stores';
import { Tag } from '@lib/types';
import cn from 'classnames';
import { Emoji, EmojiStyle } from 'emoji-picker-react';
import { Formik } from 'formik';
import { useState } from 'react';

import s from './CollectionMetadataSection.module.scss';

const MAX_NUM_TAGS = 8;

export interface ICollectionMetadataSection {
  data: {
    tags?: Tag[];
  };
  disabled?: boolean;
}

const CollectionMetadataSection: React.FC<ICollectionMetadataSection> = ({
  data,
  disabled = false,
}) => {
  const [tags, _setTags] = useState<Tag[]>(data.tags ?? []);
  const [isEditable, setIsEditable] = useState(true);
  const [isLoading, setIsLoading] = useState(false);

  const { collection, updateCollection } = useCollectionStore();

  if (!collection) return <></>;

  const initialValues: {
    title: string;
    tags: string[];
    emoji: string;
    description: string;
  } = {
    title: collection.title,
    tags: collection.tags.map((s) => s.id),
    emoji: collection.emoji,
    description: collection.description ?? '',
  };

  return (
    <Formik
      initialValues={initialValues}
      onSubmit={(values, _action) => {
        updateCollection(values);
      }}
    >
      {(props) => (
        <form
          onSubmit={props.handleSubmit}
          className={cn(
            disabled ? s.root : '',
            `${disabled ? 'w-[90%]' : 'w-full'}`
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
                <Heading as={'h3'} size={'md'} marginRight={'auto'}>
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
                    unified={collection.emoji}
                    size={35}
                    emojiStyle={EmojiStyle.TWITTER}
                  />
                  <Heading color={'white'}>{collection.title}</Heading>
                </Flex>
              )}

              {!disabled && (
                <Flex className={'!ml-auto'} alignItems={'center'} gap={5}>
                  <Box
                    display={{ base: 'none', sm: 'block' }}
                    width={{ base: '10rem', lg: '15 rem' }}
                    marginLeft={'auto'}
                  >
                    <Select
                      data={[
                        {
                          value: 'NOTE',
                          label: 'Note Collection',
                          color: 'yellow',
                        },
                        {
                          value: 'PROBLEM',
                          label: 'Problem Collection',
                          color: 'red',
                        },
                      ]}
                      value={collection.type}
                      disabled={true}
                    />
                  </Box>
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
              gridTemplateColumns={disabled ? { base: '1fr' } : { base: '1fr' }}
              gap={'5'}
              alignItems={'start'}
              justifyContent={'start'}
            >
              {!disabled && (
                <GridItem
                  width={'full'}
                  display={'grid'}
                  gridTemplateColumns={'max-content 1fr'}
                  gap={2}
                  justifyContent={'center'}
                  alignItems={'center'}
                >
                  <EmojiPicker
                    height={'full'}
                    value={props.values.emoji}
                    onChange={(emojiData, _evt) => {
                      props.setFieldValue('emoji', emojiData.unified);
                    }}
                    disabled={disabled || !isEditable}
                  />

                  <Input
                    value={props.values.title}
                    onChange={(evt) =>
                      props.setFieldValue('title', evt.target.value)
                    }
                    placeholder={'Title'}
                    readOnly={disabled || !isEditable}
                    backgroundColor={'background'}
                    borderColor={'transparent'}
                    color={'white'}
                    _hover={{
                      borderColor: 'transparent',
                    }}
                    _focusVisible={{
                      borderColor: 'primary.600',
                    }}
                  />
                </GridItem>
              )}

              <GridItem>
                {!disabled ? (
                  <Textarea
                    value={props.values.description}
                    onChange={(evt) =>
                      props.setFieldValue('description', evt.target.value)
                    }
                    placeholder={'Description'}
                    readOnly={disabled || !isEditable}
                    backgroundColor={'background'}
                    borderColor={'transparent'}
                    color={'white'}
                    _hover={{
                      borderColor: 'transparent',
                    }}
                    _focusVisible={{
                      borderColor: 'primary.600',
                    }}
                  />
                ) : (
                  <Text size={'lg'} color={'gray.300'} align={'center'}>
                    {props.values.description}
                  </Text>
                )}
              </GridItem>

              <GridItem>
                {disabled ? (
                  <Flex
                    width={'full'}
                    gap={3}
                    justifyContent={'center'}
                    flexWrap={'wrap'}
                  >
                    {collection.tags
                      .slice(0, Math.min(collection.tags.length, MAX_NUM_TAGS))
                      .map((t) => (
                        <TagLabel
                          key={t.id}
                          size={'lg'}
                          tag={t}
                          maxLabelLength={20}
                        />
                      ))}{' '}
                  </Flex>
                ) : (
                  <MultiSelect
                    data={tags.map((t) => ({
                      value: t.id,
                      label: t.label,
                      color: t.color,
                    }))}
                    value={props.values.tags}
                    onChange={(value) => props.setFieldValue('tags', value)}
                    placeholder={'Tags'}
                    disabled={disabled || !isEditable}
                  />
                )}
              </GridItem>
            </Grid>
          </Box>
        </form>
      )}
    </Formik>
  );
};

export default CollectionMetadataSection;
