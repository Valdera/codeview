import {
  Box,
  Table,
  TableContainer,
  Tbody,
  Th,
  Thead,
  Tr,
  VStack,
} from '@chakra-ui/react';
import { Problem } from '@lib/types';
import cn from 'classnames';
import { useState } from 'react';
import s from './ProblemListTable.module.scss';
import ProblemItemList from './ProblemListTableItem';
import ProblemListTablePagination from './ProblemListTablePagination';

export interface IProblemListTable {
  data?: Problem[];
  page?: number;
  totalPage?: number;
}

const ProblemListTable: React.FC<IProblemListTable> = ({
  data = [],
  page = 1,
  totalPage = 10,
}) => {
  const [problems, setProblems] = useState<Problem[]>(data);
  return (
    <VStack width={'full'} gap={2}>
      <Box
        width={'full'}
        backgroundColor={'white'}
        shadow={'sm'}
        minHeight={`${problems.length == 0 ? 'sm' : ''}`}
        rounded={'sm'}
        position={'relative'}
      >
        {problems.length > 0 ? (
          <TableContainer rounded={'sm'}>
            <Table>
              <Thead backgroundColor={'primary.700'} roundedTop={'md'}>
                <Tr>
                  <Th>
                    <h3 className={cn(s.tr)}>Title</h3>
                  </Th>
                  <Th display={{ base: 'table-cell' }}>
                    <h3 className={cn(s.tr)}>Difficulty</h3>
                  </Th>
                  <Th display={{ base: 'table-cell' }}>
                    <h3 className={cn(s.tr)}>Tags</h3>
                  </Th>
                  <Th display={{ base: 'table-cell' }}>
                    <h3 className={cn(s.tr)}>Sources</h3>
                  </Th>
                  <Th display={{ base: 'table-cell' }}>
                    <h3 className={cn(s.tr)}>Rating</h3>
                  </Th>
                </Tr>
              </Thead>
              <Tbody>
                {problems.map((problem) => (
                  <ProblemItemList key={problem.id} data={problem} />
                ))}
              </Tbody>
            </Table>
          </TableContainer>
        ) : (
          <div
            className={cn(
              'flex flex-col items-center justify-center',
              s.emptyInfo
            )}
          >
            <div
              className={
                'p-3 mx-auto text-blue-500 bg-blue-100 rounded-full dark:bg-gray-800'
              }
            >
              <svg
                xmlns={'http://www.w3.org/2000/svg'}
                fill={'none'}
                viewBox={'0 0 24 24'}
                strokeWidth={'1.5'}
                stroke={'currentColor'}
                className={'w-6 h-6'}
              >
                <path
                  strokeLinecap={'round'}
                  strokeLinejoin={'round'}
                  d={
                    'M21 21l-5.197-5.197m0 0A7.5 7.5 0 105.196 5.196a7.5 7.5 0 0010.607 10.607z'
                  }
                />
              </svg>
            </div>
            <h1 className={'mt-3 text-lg text-gray-800 dark:text-white'}>
              No problem found
            </h1>
            <p className={'mt-2 text-gray-500 dark:text-gray-400 text-center'}>
              Your search did not match any problem. Please try again or search
              another problem.
            </p>
          </div>
        )}
      </Box>
      <ProblemListTablePagination page={page} totalPage={totalPage} />
    </VStack>
  );
};

export default ProblemListTable;
