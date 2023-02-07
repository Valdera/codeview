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
import Rating from '@components/inputs/rating/Rating';
import Select from '@components/inputs/select/Select';
import { useProblemStore } from '@lib/stores';
import { Metadata } from '@lib/types/problem';
import { Formik } from 'formik';
import { useEffect, useState } from 'react';

const DIFFICULTIES_MOCK = [
  { value: 'easy_id', label: 'Easy', color: 'teal' },
  { value: 'medium_id', label: 'Medium', color: 'yellow' },
  {
    value: 'hard_id',
    label: 'Hard',
    color: 'red',
  },
];

const SOURCES_MOCK = [
  { value: 'tag_1_id', label: 'Tokopedia', color: '#FC7300' },
  { value: 'tag_2_id', label: 'Gojek', color: '#00425A' },
];

const TAGS_MOCK = [
  { value: 'tag_1_id', label: 'Array', color: '#FC7300' },
  { value: 'tag_2_id', label: 'Binary Tree', color: '#00425A' },
  { value: 'tag_3_id', label: 'Graph', color: '#BFDB38' },
  { value: 'tag_4_id', label: 'Djikstra', color: '#F55050' },
  {
    value: 'tag_5_id',
    label: 'Topological Sort',
    color: '#FF78F0',
  },
  {
    value: 'tag_6_id',
    label: 'Binary Search',
    color: '#243763',
  },
  {
    value: 'tag_7_id',
    label: 'Two Pointers',
    color: '#61876E',
  },
];

export interface IMetadataSection {}

const MetadataSection: React.FC<IMetadataSection> = () => {
  const [difficulties, setDifficulties] = useState<any>([]);
  const [sources, setSources] = useState<any>([]);
  const [tags, setTags] = useState<any>([]);

  const [isEditable, setIsEditable] = useState(true);

  const { problem, updateMetadata } = useProblemStore();

  useEffect(() => {
    setDifficulties(DIFFICULTIES_MOCK);
    setTags(TAGS_MOCK);
    setSources(SOURCES_MOCK);

    return () => {};
  }, []);

  const initialValues: Partial<Metadata> = {
    difficulty: problem.difficulty,
    tags: problem.tags,
    sources: problem.sources,
    rating: problem.rating,
  };

  return (
    <Formik
      initialValues={initialValues}
      onSubmit={(values, _actions) => {
        updateMetadata(values);
      }}
    >
      {(props) => (
        <form onSubmit={props.handleSubmit}>
          <Box width={'full'} padding={'5'} backgroundColor={'white'}>
            <HStack
              width={'full'}
              marginBottom={5}
              display={'flex'}
              alignItems={'center'}
            >
              <Heading as={'h3'} size={'md'}>
                <Highlight
                  query={'metadata'}
                  styles={{ px: '2', py: '1', rounded: 'full', bg: 'red.100' }}
                >
                  Metadata
                </Highlight>
              </Heading>
              {isEditable ? (
                <IconButton
                  className={'!ml-auto'}
                  aria-label={'Update Metadata'}
                  icon={<CheckIcon />}
                  onClick={() => {
                    props.handleSubmit();
                    setIsEditable(false);
                  }}
                />
              ) : (
                <IconButton
                  className={'!ml-auto'}
                  aria-label={'Edit Metadata'}
                  icon={<EditIcon />}
                  onClick={() => setIsEditable(true)}
                />
              )}
            </HStack>
            <Grid
              templateColumns={{
                base: '1fr 1fr',
                md: '1fr 1fr 3fr',
                lg: '1fr max-content 2fr 3fr',
              }}
              gap={'5'}
              alignItems={'start'}
              justifyContent={'start'}
            >
              <GridItem>
                <Select
                  data={difficulties}
                  value={props.values.difficulty?.id}
                  onChange={(value) => {
                    props.setFieldValue('difficulty', { id: value });
                  }}
                  label={'Difficulty'}
                  placeholder={'Select difficulty'}
                />
              </GridItem>
              <GridItem>
                <Rating
                  value={props.values.rating}
                  onChange={(value) => props.setFieldValue('rating', value)}
                  label={'Rating'}
                />
              </GridItem>
              <GridItem colSpan={{ base: 2, md: 1 }}>
                <MultiSelect
                  data={sources}
                  value={props.values.sources?.map((source) => source.id)}
                  onChange={(values) =>
                    props.setFieldValue(
                      'sources',
                      values.map((value) => ({ id: value }))
                    )
                  }
                  label={'Sources'}
                  placeholder={'Select sources'}
                />
              </GridItem>
              <GridItem colSpan={{ base: 2, md: 3, lg: 1 }}>
                <MultiSelect
                  data={tags}
                  value={props.values.tags?.map((tag) => tag.id)}
                  onChange={(values) =>
                    props.setFieldValue(
                      'tags',
                      values.map((value) => ({ id: value }))
                    )
                  }
                  label={'Tags'}
                  placeholder={'Select tags'}
                />
              </GridItem>
            </Grid>
          </Box>
        </form>
      )}
    </Formik>
  );
};

export default MetadataSection;
