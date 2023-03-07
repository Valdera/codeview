import { Box, BoxProps, Grid, GridItem } from '@chakra-ui/react';
import MultiSelect from '@components/inputs/multi-select/MultiSelect';
import Select from '@components/inputs/select/Select';
import TextInput from '@components/inputs/text/TextInput';
import { Difficulty, Source, Tag } from '@lib/types/';
import { Accordion, createStyles } from '@mantine/core';
import { IconSearch } from '@tabler/icons';
import { Formik } from 'formik';
import { ReactNode, useState } from 'react';

const useStyles = createStyles((theme) => ({
  root: {
    borderRadius: theme.radius.sm,
  },

  item: {
    backgroundColor: theme.colors.primary[7],
    border: '1px solid transparent',
    position: 'relative',
    zIndex: 0,
    transition: 'transform 150ms ease',

    '&[data-active]': {
      transform: 'scale(1.03)',
      backgroundColor:
        theme.colorScheme === 'dark' ? theme.colors.dark[7] : theme.white,
      boxShadow: theme.shadows.sm,
      borderColor:
        theme.colorScheme === 'dark'
          ? theme.colors.dark[4]
          : theme.colors.gray[2],
      borderRadius: theme.radius.md,
      zIndex: 1,
    },
  },
  control: {
    color: 'white',
    '&>*': {
      fontWeight: 500,
    },

    '&[data-active]': {
      color: theme.colors.primary[7],
    },
  },

  chevron: {
    '&[data-rotate]': {
      transform: 'rotate(-90deg)',
    },
  },
}));

export interface IProblemFilterSection extends BoxProps {
  data: {
    difficulties?: Difficulty[];
    sources?: Source[];
    tags?: Tag[];
  };

  wrapper?: 'accordion' | 'box';
}

const ProblemFilterSection: React.FC<IProblemFilterSection> = ({
  data,
  wrapper = 'box',
  ...props
}) => {
  const { classes } = useStyles();

  const [difficulties, setDifficulties] = useState<Difficulty[]>(
    data.difficulties ?? []
  );
  const [sources, setSources] = useState<Source[]>(data.sources ?? []);
  const [tags, setTags] = useState<Tag[]>(data.tags ?? []);

  const initialValues: {
    title: string;
    difficulty: string;
    tags: string[];
    sources: string[];
  } = {
    title: '',
    difficulty: '',
    tags: [],
    sources: [],
  };

  const withWrapper = (child: ReactNode): ReactNode => {
    switch (wrapper) {
      case 'accordion':
        return (
          <Box
            width={'full'}
            className={props.className}
            display={props.display}
          >
            <Accordion
              variant={'filled'}
              defaultValue={'customization'}
              classNames={classes}
              className={classes.root}
            >
              <Accordion.Item value={'customization'}>
                <Accordion.Control>Query Problem</Accordion.Control>

                <Accordion.Panel>{child}</Accordion.Panel>
              </Accordion.Item>
            </Accordion>
          </Box>
        );

      case 'box':
        return (
          <Box
            width={'full'}
            padding={'5'}
            backgroundColor={'white'}
            shadow={'sm'}
            rounded={'sm'}
            className={props.className}
            display={props.display}
          >
            {child}
          </Box>
        );
    }
  };

  return (
    <Formik
      initialValues={initialValues}
      onSubmit={(values, actions) => {
        console.log({ values, actions });
        alert(JSON.stringify(values, null, 2));
      }}
    >
      {(props) =>
        withWrapper(
          <Grid
            templateColumns={{
              base: '1fr',
              md: '1fr 1fr 1fr',
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
              />
            </GridItem>
            <GridItem className={'child:max-w-full'}>
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
              />
            </GridItem>
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
              />
            </GridItem>
            <GridItem colSpan={{ base: 1, md: 3 }} display={'grid'}>
              <TextInput
                value={props.values.title}
                onChange={(evt) =>
                  props.setFieldValue('title', evt.target.value)
                }
                label={'Title'}
                placeholder={'Search problems'}
                icon={<IconSearch size={16} />}
              />
            </GridItem>
            {/* ON CHANGE SEARCH */}
            {/* <GridItem colSpan={{ base: 1, md: 3 }}>
              <Button
                display={{ base: 'block', md: 'none' }}
                width={'full'}
                backgroundColor={'secondary.400'}
                _hover={{
                  backgroundColor: 'secondary.500',
                }}
                color={'secondary.50'}
                fontWeight={'medium'}
              >
                Search
              </Button>
            </GridItem> */}
          </Grid>
        )
      }
    </Formik>
  );
};

export default ProblemFilterSection;
