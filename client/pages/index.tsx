import { Box, Grid, GridItem, Heading } from '@chakra-ui/react';
import MultiSelect from '@components/inputs/multi-select/MultiSelect';
import Rating from '@components/inputs/rating/Rating';
import Select from '@components/inputs/select/Select';
import PrimaryLayout from '@components/layouts/primary/PrimaryLayout';
import { Formik } from 'formik';
import { useEffect, useState } from 'react';

import { NextPageWithLayout } from './page';

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

interface Metadata {
  difficulty: string;
  tags: string[];
  sources: string[];
  rating: 0 | 1 | 2 | 3 | 4 | 5;
}

const Home: NextPageWithLayout = () => {
  const [difficulties, setDifficulties] = useState<any>([]);
  const [sources, setSources] = useState<any>([]);
  const [tags, setTags] = useState<any>([]);

  useEffect(() => {
    setDifficulties(DIFFICULTIES_MOCK);
    setTags(TAGS_MOCK);
    setSources(SOURCES_MOCK);

    return () => {};
  }, []);

  const initialValues: Metadata = {
    difficulty: '',
    tags: [],
    sources: [],
    rating: 0,
  };

  return (
    <>
      <Formik
        initialValues={initialValues}
        onSubmit={(values, actions) => {
          console.log({ values, actions });
          alert(JSON.stringify(values, null, 2));
        }}
      >
        {(props) => (
          <form onSubmit={props.handleSubmit}>
            <Box width={'full'} padding={'5'} backgroundColor={'white'}>
              <Heading as={'h3'} size={'md'} marginBottom={'2'}>
                Metadata
              </Heading>
              <Grid
                templateColumns={{
                  base: '1fr 1fr',
                  md: '1fr 1fr 3fr',
                  lg: '1fr max-content 2fr 3fr',
                }}
                gap={'5'}
              >
                <GridItem>
                  <Select
                    data={difficulties}
                    value={props.values.difficulty}
                    handleChange={(value) =>
                      props.setFieldValue('difficulty', value)
                    }
                    label={'Difficulty'}
                  />
                </GridItem>
                <GridItem>
                  <Rating
                    value={props.values.rating}
                    handleChange={(value) =>
                      props.setFieldValue('rating', value)
                    }
                    label={'Rating'}
                  />
                </GridItem>
                <GridItem colSpan={{ base: 2, md: 1 }}>
                  <MultiSelect
                    data={sources}
                    value={props.values.sources}
                    handleChange={(value) =>
                      props.setFieldValue('sources', value)
                    }
                    label={'Sources'}
                  />
                </GridItem>
                <GridItem colSpan={{ base: 2, md: 3, lg: 1 }}>
                  <MultiSelect
                    data={tags}
                    value={props.values.tags}
                    handleChange={(value) => props.setFieldValue('tags', value)}
                    label={'Tags'}
                  />
                </GridItem>
              </Grid>
            </Box>
          </form>
        )}
      </Formik>
    </>
  );
};

export default Home;

Home.getLayout = (page) => {
  return <PrimaryLayout justify={'items-start'}>{page}</PrimaryLayout>;
};
