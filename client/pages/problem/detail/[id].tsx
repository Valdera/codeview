import { VStack } from '@chakra-ui/react';
import PrimaryLayout from '@components/layouts/primary/PrimaryLayout';
import ProblemMetadataSection from '@components/problem/sections/metadata/ProblemMetadataSection';
import ProblemQuestionSection from '@components/problem/sections/question/ProblemQuestionSection';
import ProblemSolutionSection from '@components/problem/sections/solution/ProblemSolutionSection';
import { useProblemStore } from '@lib/stores';
import { Difficulty, Problem, Source, Tag } from '@lib/types';
import { NextPageWithLayout } from '@pages/page';
import { GetServerSideProps, InferGetServerSidePropsType } from 'next';
import { useEffect } from 'react';

interface IProblemDetailPage {
  problem: Problem;
  difficulties: Difficulty[];
  sources: Source[];
  tags: Tag[];
}

const ProblemDetailPage: NextPageWithLayout<IProblemDetailPage> = ({
  problem: initialProblem,
  difficulties,
  sources,
  tags,
}: InferGetServerSidePropsType<typeof getServerSideProps>) => {
  const { problem, load, createSolution } = useProblemStore();

  useEffect(() => {
    load(initialProblem);

    return () => {};
  }, [load, initialProblem]);

  return (
    <>
      {problem && problem.question && (
        <VStack width={'full'} gap={5}>
          <ProblemMetadataSection
            data={{ difficulties, sources, tags }}
            disabled={true}
          />
          <ProblemQuestionSection question={problem.question} disabled={true} />
          {problem.solutions?.map((solution) => (
            <ProblemSolutionSection
              key={solution.id}
              solution={solution}
              disabled={true}
            />
          ))}
        </VStack>
      )}
    </>
  );
};

ProblemDetailPage.getLayout = (page) => {
  return <PrimaryLayout justify={'items-start'}>{page}</PrimaryLayout>;
};

export const getServerSideProps: GetServerSideProps<
  IProblemDetailPage
> = async (_context) => {
  const MOCK_PROBLEM: Problem = {
    id: '1',
    title: 'Median in Array Streams',
    rating: 3,
    difficulty: { id: '1', label: 'Easy', color: 'teal' },
    sources: [{ id: '1', label: 'Tokopedia', color: '#FC7300' }],
    tags: [
      { id: '2', label: 'Binary Tree', color: '#00425A' },
      { id: '3', label: 'Graph', color: '#BFDB38' },
      { id: '4', label: 'Djikstra', color: '#F55050' },
    ],
    question: {
      id: '1',
      content: `<h2 style="text-align: center; margin-left: 0px!important;"><strong>Problem Question</strong></h2>`,
    },
    solutions: [
      {
        id: '1',
        content: `<h2 style="text-align: center; margin-left: 0px!important;"><strong>Problem Solution #1</strong></h2>`,
      },
      {
        id: '2',
        content: `<h2 style="text-align: center; margin-left: 0px!important;"><strong>Problem Solution #2</strong></h2>`,
      },
    ],
  };

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

  return {
    props: {
      difficulties: MOCK_DIFFICULTIES,
      tags: MOCK_TAGS,
      sources: MOCK_SOURCES,
      problem: MOCK_PROBLEM,
    },
  };
};

export default ProblemDetailPage;
