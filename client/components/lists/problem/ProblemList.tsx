import { Table, TableContainer, Tbody, Th, Thead, Tr } from '@chakra-ui/react';
import { useEffect, useState } from 'react';
import ProblemItemList from './ProblemItemList';

const PROBLEMS_MOCK = [
  {
    title: 'Median of Two Sorted Arrays',
    slug: '/problems/median-of-two-sorted-arrays',
    difficulty: { label: 'easy', color: 'teal' },
    tags: [
      { label: 'Array', color: '#FC7300' },
      { label: 'Binary Tree', color: '#00425A' },
      { label: 'Graph', color: '#BFDB38' },
    ],
    sources: [
      { label: 'Tokopedia', color: '#FC7300' },
      { label: 'Gojek', color: '#00425A' },
    ],
    rating: 4,
  },
  {
    title: 'Median of Two Sorted Arrays',
    slug: '/problems/median-of-two-sorted-arrays',
    difficulty: { label: 'easy', color: 'teal' },
    tags: [
      { label: 'Array', color: '#FC7300' },
      { label: 'Binary Tree', color: '#00425A' },
      { label: 'Graph', color: '#BFDB38' },
    ],
    sources: [
      { label: 'Tokopedia', color: '#FC7300' },
      { label: 'Gojek', color: '#00425A' },
    ],
    rating: 4,
  },
  {
    title: 'Median of Two Sorted Arrays',
    slug: '/problems/median-of-two-sorted-arrays',
    difficulty: { label: 'easy', color: 'teal' },
    tags: [
      { label: 'Array', color: '#FC7300' },
      { label: 'Binary Tree', color: '#00425A' },
      { label: 'Graph', color: '#BFDB38' },
    ],
    sources: [
      { label: 'Tokopedia', color: '#FC7300' },
      { label: 'Gojek', color: '#00425A' },
    ],
    rating: 4,
  },
  {
    title: 'Median of Two Sorted Arrays',
    slug: '/problems/median-of-two-sorted-arrays',
    difficulty: { label: 'easy', color: 'teal' },
    tags: [
      { label: 'Array', color: '#FC7300' },
      { label: 'Binary Tree', color: '#00425A' },
      { label: 'Graph', color: '#BFDB38' },
    ],
    sources: [
      { label: 'Tokopedia', color: '#FC7300' },
      { label: 'Gojek', color: '#00425A' },
    ],
    rating: 4,
  },
];

export interface IProblemList {}

const ProblemList: React.FC<IProblemList> = () => {
  const [problems, setProblems] = useState<any>([]);

  useEffect(() => {
    setProblems(PROBLEMS_MOCK);

    return () => {};
  }, []);

  return (
    <TableContainer>
      <Table variant={'striped'}>
        <Thead>
          <Tr>
            <Th>Title</Th>
            <Th display={{ base: 'none', sm: 'table-cell' }}>Difficulty</Th>
            <Th display={{ base: 'none', md: 'table-cell' }}>Tags</Th>
            <Th display={{ base: 'none', lg: 'table-cell' }}>Sources</Th>
            <Th display={{ base: 'none', '2xl': 'table-cell' }}>Rating</Th>
          </Tr>
        </Thead>
        <Tbody>
          {problems.map((problem: any) => (
            <ProblemItemList
              key={problem.title}
              slug={problem.slug}
              title={problem.title}
              difficulty={problem.difficulty}
              tags={problem.tags}
              sources={problem.sources}
              rating={problem.rating}
            />
          ))}
        </Tbody>
      </Table>
    </TableContainer>
  );
};

export default ProblemList;
