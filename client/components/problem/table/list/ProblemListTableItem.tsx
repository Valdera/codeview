import { Grid, Tag, Td, Tr } from '@chakra-ui/react';
import Rating from '@components/inputs/rating/Rating';
import { Problem } from '@lib/types';
import Link from 'next/link';

export interface IProblemListTableItem {
  data: Problem;
}

const ProblemListTableItem: React.FC<IProblemListTableItem> = ({ data }) => {
  return (
    <Tr>
      <Td>
        <Link
          href={data.slug ?? '/'}
          className={'text-blue-400 hover:text-blue-600 transition-all'}
        >
          {data.title}
        </Link>
      </Td>
      <Td display={{ base: 'table-cell' }}>
        <Tag
          colorScheme={data.difficulty.color}
          width={'100px'}
          display={'flex'}
          justifyContent={'center'}
          fontWeight={'bold'}
        >
          {data.difficulty.label.toUpperCase()}
        </Tag>
      </Td>
      <Td display={{ base: 'table-cell' }}>
        <Grid
          display={'grid'}
          gap={2}
          gridTemplateColumns={{
            base: 'repeat(1, max-content)',
            lg: 'repeat(2, max-content)',
            xl: 'repeat(3, max-content)',
          }}
        >
          {data.tags.map((tag) => (
            <Tag
              key={tag.label}
              width={'100px'}
              display={'flex'}
              backgroundColor={tag.color}
              color={'white'}
              justifyContent={'center'}
            >
              {tag.label}
            </Tag>
          ))}
        </Grid>
      </Td>
      <Td display={{ base: 'table-cell' }}>
        <Grid
          display={'grid'}
          gap={2}
          gridTemplateColumns={{
            base: 'repeat(1, max-content)',
            xl: 'repeat(2, max-content)',
          }}
        >
          {data.sources.map((source) => (
            <Tag
              key={source.label}
              width={'100px'}
              display={'flex'}
              backgroundColor={source.color}
              color={'white'}
              justifyContent={'center'}
            >
              {source.label}
            </Tag>
          ))}
        </Grid>
      </Td>
      <Td display={{ base: 'table-cell' }}>
        <Rating value={data.rating} disabled={true} />
      </Td>
    </Tr>
  );
};

export default ProblemListTableItem;
