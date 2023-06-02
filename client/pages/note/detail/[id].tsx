import { VStack } from '@chakra-ui/react';
import PrimaryLayout from '@components/layouts/primary/PrimaryLayout';
import NoteItemSection from '@components/note/sections/item/NoteItemSection';
import NoteMetadataSection from '@components/note/sections/metadata/NoteMetadataSection';
import { useNoteStore } from '@lib/stores';
import { Note } from '@lib/types';
import { NextPageWithLayout } from '@pages/page';
import { GetServerSidePropsContext, GetServerSidePropsResult } from 'next';
import { useEffect } from 'react';

interface INoteEditPage {
  note: Note;
}

const NoteEditPage: NextPageWithLayout<INoteEditPage> = ({
  note: initialNote,
}) => {
  const { load, note } = useNoteStore();

  useEffect(() => {
    load(initialNote);
  }, [load, initialNote]);

  if (!note) return <></>;

  return (
    <>
      <VStack width={'full'} gap={5}>
        <NoteMetadataSection data={{}} disabled={true} />
        <NoteItemSection disabled={true} />
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
        content: `<h2 style="text-align: center; margin-left: 0px!important;"><strong>The Problem of Granularity</strong></h2>`,
        numOrder: 0,
      },
      {
        id: '2',
        content: `<h2 style="text-align: center; margin-left: 0px!important;"><strong>The Problem of Identity</strong></h2>`,
        numOrder: 1,
      },
      {
        id: '3',
        content: `<h2 style="text-align: center; margin-left: 0px!important;"><strong>The Problem of Inheritance</strong></h2>`,
        numOrder: 2,
      },
    ],
    status: 'DRAFT',
    emoji: '1f92a',
  };

  return {
    props: {
      note: MOCK_NOTE,
    },
  };
};

export default NoteEditPage;
