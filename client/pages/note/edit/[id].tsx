import { AddIcon } from '@chakra-ui/icons';
import { Button, VStack } from '@chakra-ui/react';
import PrimaryLayout from '@components/layouts/primary/PrimaryLayout';
import NoteItemSection from '@components/note/sections/item/NoteItemSection';
import NoteMetadataSection from '@components/note/sections/metadata/NoteMetadataSection';
import { useNoteStore } from '@lib/stores';
import { Note, Tag } from '@lib/types';
import { NextPageWithLayout } from '@pages/page';
import { GetServerSidePropsContext, GetServerSidePropsResult } from 'next';
import { useEffect } from 'react';
import { DragDropContext, Droppable } from 'react-beautiful-dnd';

interface INoteEditPage {
  note: Note;
  tags: Tag[];
}

const NoteEditPage: NextPageWithLayout<INoteEditPage> = ({
  tags,
  note: initialNote,
}) => {
  const { reorderNoteItem, load, note, createNoteItem } = useNoteStore();

  useEffect(() => {
    load(initialNote);
  }, [load, initialNote]);

  if (!note) return <></>;

  return (
    <>
      <VStack width={'full'} gap={5}>
        <NoteMetadataSection data={{ note: note, tags: tags }} />
        <DragDropContext
          onDragEnd={({ destination, source }) =>
            reorderNoteItem(source.index, destination?.index || 0)
          }
        >
          <Droppable droppableId={'dnd-list'} direction={'vertical'}>
            {(provided) => (
              <div
                className={'w-full flex flex-col'}
                {...provided.droppableProps}
                ref={provided.innerRef}
              >
                {note.items?.map((item, index) => (
                  <NoteItemSection
                    key={item.id}
                    index={index}
                    data={{ noteItem: item }}
                  />
                ))}
                {provided.placeholder}
              </div>
            )}
          </Droppable>
        </DragDropContext>
        <Button
          onClick={() => createNoteItem()}
          rightIcon={<AddIcon />}
          colorScheme={'purple'}
          width={'full'}
        >
          Add Note Item
        </Button>
      </VStack>
    </>
  );
};

NoteEditPage.getLayout = (page) => {
  return <PrimaryLayout justify={'items-start'}>{page}</PrimaryLayout>;
};

export const getServerSideProps = async (
  _context: GetServerSidePropsContext
): Promise<GetServerSidePropsResult<INoteEditPage>> => {
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

  const MOCK_NOTE: Note = {
    id: '1',
    title: 'The Paradigm Mismatch',
    tags: [
      { id: '1', label: 'Spring Boot', color: '#FC7300' },
      { id: '3', label: 'Java', color: '#BFDB38' },
    ],
    items: [
      {
        id: '1',
        title: 'The Problem of Granularity',
        content: `<h2 style="text-align: center; margin-left: 0px!important;"><strong>The Problem of Granularity</strong></h2>`,
        numOrder: 0,
      },
      {
        id: '2',
        title: 'The Problem of Identity',
        content: `<h2 style="text-align: center; margin-left: 0px!important;"><strong>The Problem of Identity</strong></h2>`,
        numOrder: 1,
      },
      {
        id: '3',
        title: 'The Problem of Inheritance',
        content: `<h2 style="text-align: center; margin-left: 0px!important;"><strong>The Problem of Inheritance</strong></h2>`,
        numOrder: 2,
      },
    ],
    references: ['Java Persistence by Oreilly'],
  };

  return {
    props: {
      note: MOCK_NOTE,
      tags: MOCK_TAGS,
    },
  };
};

export default NoteEditPage;
