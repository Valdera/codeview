import { Grid, Tag, Td, Tr } from '@chakra-ui/react';
import Rating from '@components/inputs/rating/Rating';
import Link from 'next/link';

export interface IProblemItemList {
  title: string;
  slug: string;
  difficulty: { label: string; color: string };
  tags: { label: string; color: string }[];
  sources: { label: string; color: string }[];
  rating: 0 | 1 | 2 | 3 | 4 | 5;
}

const ProblemItemList: React.FC<IProblemItemList> = ({
  title,
  slug,
  difficulty,
  tags,
  sources,
  rating,
}) => {
  return (
    <Tr>
      <Td>
        <Link
          href={slug}
          className={'text-blue-400 hover:text-blue-600 transition-all'}
        >
          {title}
        </Link>
      </Td>
      <Td display={{ base: 'none', sm: 'table-cell' }}>
        <Tag
          colorScheme={difficulty.color}
          width={'100px'}
          display={'flex'}
          justifyContent={'center'}
        >
          {difficulty.label}
        </Tag>
      </Td>
      <Td display={{ base: 'none', md: 'table-cell' }}>
        <Grid
          display={'grid'}
          gap={2}
          gridTemplateColumns={{
            base: 'repeat(1, max-content)',
            lg: 'repeat(2, max-content)',
            xl: 'repeat(3, max-content)',
          }}
        >
          {tags.map((tag) => (
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
      <Td display={{ base: 'none', lg: 'table-cell' }}>
        <Grid
          display={'grid'}
          gap={2}
          gridTemplateColumns={{
            base: 'repeat(1, max-content)',
            xl: 'repeat(2, max-content)',
          }}
        >
          {sources.map((source) => (
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
      <Td display={{ base: 'none', '2xl': 'table-cell' }}>
        <Rating value={rating} disabled={true} />
      </Td>
    </Tr>
  );
};

export default ProblemItemList;
