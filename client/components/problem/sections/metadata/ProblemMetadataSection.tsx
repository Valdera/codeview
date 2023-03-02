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
import { Difficulty, Source, Tag } from '@lib/types';
import { Formik } from 'formik';
import { useState } from 'react';

export interface IProblemMetadataSection {
  data: {
    difficulties?: Difficulty[];
    sources?: Source[];
    tags?: Tag[];
  };
  disabled?: boolean;
}

const ProblemMetadataSection: React.FC<IProblemMetadataSection> = ({
  data,
  disabled = false,
}) => {
  const [difficulties, setDifficulties] = useState<Difficulty[]>(
    data.difficulties ?? []
  );
  const [sources, setSources] = useState<Source[]>(data.sources ?? []);
  const [tags, setTags] = useState<Tag[]>(data.tags ?? []);

  const [isEditable, setIsEditable] = useState(true);
  const [isLoading, setIsLoading] = useState(false);

  const { problem } = useProblemStore();
  if (!problem) return <></>;

  const initialValues: {
    difficulty: string;
    tags: string[];
    sources: string[];
    rating: 0 | 1 | 2 | 3 | 4 | 5;
  } = {
    difficulty: problem.difficulty.id,
    tags: problem.tags.map((s) => s.id),
    sources: problem.sources.map((s) => s.id),
    rating: problem.rating,
  };

  return (
    <Formik
      initialValues={initialValues}
      onSubmit={(values, _actions) => {
        console.log(values);
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
                  data={difficulties.map((d) => ({
                    value: d.id,
                    label: d.label,
                    color: d.color,
                  }))}
                  value={props.values.difficulty}
                  onChange={(value) => props.setFieldValue('difficulty', value)}
                  label={'Difficulty'}
                  placeholder={'Select difficulty'}
                  disabled={!isEditable}
                />
              </GridItem>
              <GridItem>
                <Rating
                  value={props.values.rating}
                  onChange={(value) => props.setFieldValue('rating', value)}
                  label={'Rating'}
                  disabled={!isEditable}
                />
              </GridItem>
              <GridItem colSpan={{ base: 2, md: 1 }}>
                <MultiSelect
                  data={sources.map((s) => ({
                    value: s.id,
                    label: s.label,
                    color: s.color,
                  }))}
                  value={props.values.sources}
                  onChange={(value) => props.setFieldValue('sources', value)}
                  label={'Sources'}
                  placeholder={'Select sources'}
                  disabled={!isEditable}
                />
              </GridItem>
              <GridItem colSpan={{ base: 2, md: 3, lg: 1 }}>
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
                  disabled={!isEditable}
                />
              </GridItem>
            </Grid>
          </Box>
        </form>
      )}
    </Formik>
  );
};

export default ProblemMetadataSection;
