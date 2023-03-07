import { Grid, Tag, Td, Tr } from '@chakra-ui/react';
import { Note } from '@lib/types';
import Link from 'next/link';

export interface INoteListTableItem {
  data: Note;
}

const NoteListTableItem: React.FC<INoteListTableItem> = ({ data }) => {
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
    </Tr>
  );
};

export default NoteListTableItem;
