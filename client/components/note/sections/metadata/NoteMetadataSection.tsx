import { CheckIcon, EditIcon } from '@chakra-ui/icons';
import {
  Box,
  Grid,
  GridItem,
  Heading,
  Highlight,
  HStack,
  IconButton,
} from '@chakra-ui/react';
import MultiSelect from '@components/inputs/multi-select/MultiSelect';
import TextInput from '@components/inputs/text/TextInput';
import { Note, Tag } from '@lib/types';
import { Formik } from 'formik';
import { useState } from 'react';

export interface INoteMetadataSection {
  data: {
    note: Note;
    tags?: Tag[];
  };
  disabled?: boolean;
}

const NoteMetadataSection: React.FC<INoteMetadataSection> = ({
  data,
  disabled = false,
}) => {
  const [tags, setTags] = useState<Tag[]>(data.tags ?? []);
  const [isEditable, setIsEditable] = useState(true);
  const [isLoading, setIsLoading] = useState(false);

  const note = data.note;

  const initialValues: {
    title: string;
    tags: string[];
  } = {
    title: note.title,
    tags: note.tags.map((s) => s.id),
  };

  return (
    <Formik
      initialValues={initialValues}
      onSubmit={(values, _action) => {
        console.log(values);
      }}
    >
      {(props) => (
        <form onSubmit={props.handleSubmit} className={'w-full'}>
          <Box width={'full'} padding={'5'} backgroundColor={'white'}>
            <HStack
              width={'full'}
              marginBottom={5}
              display={'flex'}
              alignItems={'center'}
            >
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
                <Heading>{note.title}</Heading>
              )}
              {!disabled && (
                <>
                  {isEditable ? (
                    <IconButton
                      className={'!ml-auto'}
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
                </>
              )}
            </HStack>
            <Grid
              gridTemplateColumns={
                disabled
                  ? { base: '1fr', md: '400px' }
                  : { base: '1fr', md: '3fr 1fr' }
              }
              gap={'5'}
              alignItems={'start'}
              justifyContent={'start'}
            >
              {!disabled && (
                <GridItem>
                  <TextInput
                    value={props.values.title}
                    onChange={(evt) =>
                      props.setFieldValue('title', evt.target.value)
                    }
                    label={'Title'}
                    placeholder={'Search problems'}
                  />
                </GridItem>
              )}
              <GridItem>
                <MultiSelect
                  data={tags.map((t) => ({
                    value: t.id,
                    label: t.label,
                    color: t.color,
                  }))}
                  value={props.values.tags}
                  onChange={(value) => props.setFieldValue('tags', value)}
                  label={'Tags'}
                  placeholder={'Select tags'}
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
