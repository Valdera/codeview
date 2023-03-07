import { Grid, useBreakpointValue, VStack } from '@chakra-ui/react';
import PreviewCarousel from '@components/carousels/preview/PreviewCarousel';
import PrimaryLayout from '@components/layouts/primary/PrimaryLayout';
import ProblemFilterSection from '@components/problem/sections/filter/ProblemFilterSection';
import ProblemListTable from '@components/problem/table/list/ProblemListTable';
import SegmentStat from '@components/stats/segment/SegmentStat';
import { Difficulty, Problem, Source, Tag } from '@lib/types';
import { NextPageWithLayout } from '@pages/page';
import {
  GetServerSidePropsContext,
  GetServerSidePropsResult,
  InferGetServerSidePropsType,
} from 'next';

interface IProblemListPage {
  difficulties: Difficulty[];
  tags: Tag[];
  sources: Source[];
  problems: Problem[];
}

const ProblemListPage: NextPageWithLayout<IProblemListPage> = ({
  difficulties,
  sources,
  tags,
  problems,
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
      <ProblemFilterSection
        wrapper={filterWrapper}
        data={{ difficulties, sources, tags }}
      />
      <ProblemListTable data={problems} />
    </VStack>
  );
};

ProblemListPage.getLayout = (page) => {
  return <PrimaryLayout justify={'items-start'}>{page}</PrimaryLayout>;
};

export const getServerSideProps = async (
  _context: GetServerSidePropsContext
): Promise<GetServerSidePropsResult<IProblemListPage>> => {
  const MOCK_DIFFICULTIES: Difficulty[] = [
    { id: '1', label: 'Easy', color: 'teal' },
    { id: '2', label: 'Medium', color: 'yellow' },
    {
      id: '3',
      label: 'Hard',
      color: 'red',
    },
  ];

  const MOCK_SOURCES: Source[] = [
    { id: '1', label: 'Tokopedia', color: '#FC7300' },
    { id: '2', label: 'Gojek', color: '#00425A' },
  ];

  const MOCK_TAGS: Tag[] = [
    { id: '1', label: 'Array', color: '#FC7300' },
    { id: '2', label: 'Binary Tree', color: '#00425A' },
    { id: '3', label: 'Graph', color: '#BFDB38' },
    { id: '4', label: 'Djikstra', color: '#F55050' },
    {
      id: '5',
      label: 'Topological Sort',
      color: '#FF78F0',
    },
    {
      id: '6',
      label: 'Binary Search',
      color: '#243763',
    },
    {
      id: '7',
      label: 'Two Pointers',
      color: '#61876E',
    },
  ];

  const MOCK_PROBLEMS: Problem[] = [
    {
      id: '1',
      title: 'Median of Two Sorted Arrays',
      slug: '/problem/detail/1',
      difficulty: { id: '1', label: 'easy', color: 'teal' },
      tags: [
        { id: '1', label: 'Array', color: '#FC7300' },
        { id: '2', label: 'Binary Tree', color: '#00425A' },
        { id: '3', label: 'Graph', color: '#BFDB38' },
      ],
      sources: [
        { id: '1', label: 'Tokopedia', color: '#FC7300' },
        { id: '2', label: 'Gojek', color: '#00425A' },
      ],
      rating: 4,
    },
    {
      id: '2',
      title: 'Median of Two Sorted Arrays',
      slug: '/problem/detail/1',
      difficulty: { id: '1', label: 'easy', color: 'teal' },
      tags: [
        { id: '1', label: 'Array', color: '#FC7300' },
        { id: '2', label: 'Binary Tree', color: '#00425A' },
        { id: '3', label: 'Graph', color: '#BFDB38' },
      ],
      sources: [
        { id: '1', label: 'Tokopedia', color: '#FC7300' },
        { id: '2', label: 'Gojek', color: '#00425A' },
      ],
      rating: 4,
    },
    {
      id: '3',
      title: 'Median of Two Sorted Arrays',
      slug: '/problem/detail/1',
      difficulty: { id: '1', label: 'easy', color: 'teal' },
      tags: [
        { id: '1', label: 'Array', color: '#FC7300' },
        { id: '2', label: 'Binary Tree', color: '#00425A' },
        { id: '3', label: 'Graph', color: '#BFDB38' },
      ],
      sources: [
        { id: '1', label: 'Tokopedia', color: '#FC7300' },
        { id: '2', label: 'Gojek', color: '#00425A' },
      ],
      rating: 4,
    },
    {
      id: '4',
      title: 'Median of Two Sorted Arrays',
      slug: '/problem/detail/1',
      difficulty: { id: '1', label: 'easy', color: 'teal' },
      tags: [
        { id: '1', label: 'Array', color: '#FC7300' },
        { id: '2', label: 'Binary Tree', color: '#00425A' },
        { id: '3', label: 'Graph', color: '#BFDB38' },
      ],
      sources: [
        { id: '1', label: 'Tokopedia', color: '#FC7300' },
        { id: '2', label: 'Gojek', color: '#00425A' },
      ],
      rating: 4,
    },
  ];

  return {
    props: {
      difficulties: MOCK_DIFFICULTIES,
      tags: MOCK_TAGS,
      sources: MOCK_SOURCES,
      problems: MOCK_PROBLEMS,
    },
  };
};

export default ProblemListPage;
