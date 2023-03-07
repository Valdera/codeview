import { Grid, useBreakpointValue, VStack } from '@chakra-ui/react';
import PreviewCarousel from '@components/carousels/preview/PreviewCarousel';
import PrimaryLayout from '@components/layouts/primary/PrimaryLayout';
import NoteListTable from '@components/note/table/list/NoteListTable';
import SegmentStat from '@components/stats/segment/SegmentStat';
import { Note, Tag } from '@lib/types';
import { NextPageWithLayout } from '@pages/page';
import {
  GetServerSidePropsContext,
  GetServerSidePropsResult,
  InferGetServerSidePropsType,
} from 'next';

interface IProblemListPage {
  notes: Note[];
  tags: Tag[];
}

const ProblemListPage: NextPageWithLayout<IProblemListPage> = ({
  notes,
  tags,
}: InferGetServerSidePropsType<typeof getServerSideProps>) => {
  const filterWrapper: 'accordion' | 'box' =
    useBreakpointValue(
      {
        base: 'accordion',
        md: 'box',
      },
      {
        ssr: true,
        fallback: 'box',
      }
    ) ?? 'box';

  return (
    <VStack width={'full'} gap={5}>
      <Grid
        width={'full'}
        gridTemplateColumns={{ base: '1fr', lg: 'max-content 1fr' }}
        gridTemplateRows={{ base: 'max-content 200px', lg: 'max-content' }}
        gap={5}
      >
        <SegmentStat
          total={'200 Question'}
          data={[
            {
              label: 'Easy Problems',
              count: '10',
              part: 30,
              color: '#4FD1C5',
            },
            {
              label: 'Medium Problems',
              count: '10',
              part: 30,
              color: '#F6E05E',
            },
            {
              label: 'Hard Problems',
              count: '10',
              part: 40,
              color: '#F56565',
            },
          ]}
        />
        <PreviewCarousel />
      </Grid>
      <NoteListTable data={notes} />
    </VStack>
  );
};

ProblemListPage.getLayout = (page) => {
  return <PrimaryLayout justify={'items-start'}>{page}</PrimaryLayout>;
};

export const getServerSideProps = async (
  _context: GetServerSidePropsContext
): Promise<GetServerSidePropsResult<IProblemListPage>> => {
  const MOCK_TAGS: Tag[] = [
    { id: '1', label: 'Spring Boot', color: '#FC7300' },
    { id: '2', label: 'Go', color: '#00425A' },
    { id: '3', label: 'Java', color: '#BFDB38' },
    { id: '4', label: 'Micronaut', color: '#F55050' },
    {
      id: '5',
      label: 'Clean Architecture',
      color: '#FF78F0',
    },
    {
      id: '6',
      label: 'Clojure',
      color: '#243763',
    },
  ];

  const MOCK_NOTES: Note[] = [
    {
      id: '1',
      title: 'The Paradigm Mismatch',
      slug: '/problem/detail/1',
      tags: [
        { id: '1', label: 'Spring Boot', color: '#FC7300' },
        { id: '3', label: 'Java', color: '#BFDB38' },
      ],
      references: ['Java Persistence by Oreilly'],
    },
    {
      id: '2',
      title: 'The Paradigm Mismatch',
      slug: '/problem/detail/2',
      tags: [
        { id: '1', label: 'Spring Boot', color: '#FC7300' },
        { id: '3', label: 'Java', color: '#BFDB38' },
      ],
      references: ['Java Persistence by Oreilly'],
    },
    {
      id: '3',
      title: 'Micronaut Hexa Architecture',
      slug: '/problem/detail/3',
      tags: [
        { id: '4', label: 'Micronaut', color: '#F55050' },
        {
          id: '5',
          label: 'Clean Architecture',
          color: '#FF78F0',
        },
      ],
      references: ['Java Persistence by Oreilly'],
    },
  ];

  return {
    props: {
      notes: MOCK_NOTES,
      tags: MOCK_TAGS,
    },
  };
};

export default ProblemListPage;
