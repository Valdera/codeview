import { Box, Grid, GridItem, VStack } from '@chakra-ui/react';
import MultiSelect from '@components/inputs/multi-select/MultiSelect';
import Select from '@components/inputs/select/Select';
import TextInput from '@components/inputs/text-input/TextInput';
import PrimaryLayout from '@components/layouts/primary/PrimaryLayout';
import ProblemList from '@components/lists/problem/ProblemList';
import { NextPageWithLayout } from '@pages/page';
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

interface ProblemMetadata {
  title: string;
  difficulty: string;
  tags: string[];
  sources: string[];
}

const ProblemsPage: NextPageWithLayout = () => {
  const [difficulties, setDifficulties] = useState<any>([]);
  const [sources, setSources] = useState<any>([]);
  const [tags, setTags] = useState<any>([]);

  useEffect(() => {
    setDifficulties(DIFFICULTIES_MOCK);
    setTags(TAGS_MOCK);
    setSources(SOURCES_MOCK);

    return () => {};
  }, []);

  const initialValues: ProblemMetadata = {
    title: '',
    difficulty: '',
    tags: [],
    sources: [],
  };

  return (
    <VStack width={'full'}>
      <Formik
        initialValues={initialValues}
        onSubmit={(values, actions) => {
          console.log({ values, actions });
          alert(JSON.stringify(values, null, 2));
        }}
      >
        {(props) => (
          <Box
            width={'full'}
            padding={'5'}
            backgroundColor={'white'}
            shadow={'sm'}
            rounded={'md'}
          >
            <Grid
              templateColumns={{
                base: '1fr 1fr 1fr 1fr',
              }}
              gap={'5'}
              alignItems={'start'}
              justifyContent={'start'}
            >
              <GridItem>
                <Select
                  data={difficulties}
                  value={props.values.difficulty}
                  onChange={(value) => props.setFieldValue('difficulty', value)}
                  label={'Difficulty'}
                  placeholder={'Select difficulty'}
                />
              </GridItem>

              <GridItem
                colSpan={{ base: 2, md: 1 }}
                className={'child:max-w-full'}
              >
                <MultiSelect
                  data={sources}
                  value={props.values.sources}
                  onChange={(value) => props.setFieldValue('sources', value)}
                  label={'Sources'}
                  placeholder={'Select sources'}
                />
              </GridItem>
              <GridItem colSpan={{ base: 2, md: 3, lg: 1 }}>
                <MultiSelect
                  data={tags}
                  value={props.values.tags}
                  onChange={(value) => props.setFieldValue('tags', value)}
                  label={'Tags'}
                  placeholder={'Select tags'}
                />
              </GridItem>
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
            </Grid>
          </Box>
        )}
      </Formik>
      <Box
        width={'full'}
        padding={'5'}
        backgroundColor={'white'}
        shadow={'sm'}
        rounded={'md'}
      >
        <ProblemList />
      </Box>
    </VStack>
  );
};

export default ProblemsPage;

ProblemsPage.getLayout = (page) => {
  return <PrimaryLayout justify={'items-start'}>{page}</PrimaryLayout>;
};
