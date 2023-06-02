import { ChevronRightIcon } from '@chakra-ui/icons';
import { Flex, Grid, Text, VStack } from '@chakra-ui/react';
import PreviewCarousel from '@components/carousels/preview/PreviewCarousel';
import CollectionCarouselCard from '@components/collection/cards/carousel/CollectionCarouselCard';
import SearchBar from '@components/inputs/search/SearchBar';
import PrimaryLayout from '@components/layouts/primary/PrimaryLayout';
import NoteListCard from '@components/note/cards/list/NoteListCard';
import { CollectionPreview, Note } from '@lib/types';
import { NextPageWithLayout } from '@pages/page';
import {
  GetServerSidePropsContext,
  GetServerSidePropsResult,
  InferGetServerSidePropsType,
} from 'next';

interface INoteListPage {
  notes: Note[];
  collectionPreview: CollectionPreview[];
}

const NoteListPage: NextPageWithLayout<INoteListPage> = ({
  notes,
  collectionPreview,
}: InferGetServerSidePropsType<typeof getServerSideProps>) => {
  return (
    <VStack width={'full'} gap={5}>
      <Grid
        width={'full'}
        gridTemplateColumns={{ base: '1fr', lg: 'max-content 1fr' }}
        gridTemplateRows={{ base: 'max-content 200px', lg: 'max-content' }}
        gap={5}
        alignItems={'center'}
      >
        <Flex
          padding={'5'}
          backgroundColor={'foreground'}
          shadow={'md'}
          flexDir={{ base: 'row', lg: 'column' }}
          borderRadius={'md'}
          fontWeight={'semibold'}
          cursor={'pointer'}
          transition={'all .5s'}
          _hover={{ backgroundColor: 'primary.600' }}
        >
          <Text
            noOfLines={2}
            fontFamily={'heading'}
            color={'white'}
            marginBottom={'2'}
            height={'70px'}
            fontSize={'3xl'}
          >
            Note Collections
          </Text>
          <ChevronRightIcon
            color={'white'}
            fontSize={'5xl'}
            marginLeft={{ base: 'auto', lg: '0' }}
            alignSelf={{ base: 'center', lg: 'start' }}
          />
        </Flex>
        <Grid width={'full'} height={'90%'}>
          <PreviewCarousel
            contents={collectionPreview.map((item) => (
              <CollectionCarouselCard key={item.id} {...item} />
            ))}
          />
        </Grid>
      </Grid>
      <Flex alignItems={'center'} justifyContent={'center'} width={'full'}>
        <Grid
          width={{ base: 'full', lg: '80%' }}
          borderRadius={'xl'}
          padding={'5'}
          backgroundColor={'foreground'}
        >
          <SearchBar />
        </Grid>
      </Flex>

      <Grid
        width={'full'}
        gridTemplateColumns={{
          base: 'repeat(2, 1fr)',
          lg: 'repeat(3, 1fr)',
          xl: 'repeat(4, 1fr)',
          '2xl': 'repeat(5, 1fr)',
        }}
        gap={5}
      >
        {notes.map((note) => (
          <NoteListCard key={note.id} note={note} />
        ))}
      </Grid>
    </VStack>
  );
};

NoteListPage.getLayout = (page) => {
  return <PrimaryLayout justify={'items-start'}>{page}</PrimaryLayout>;
};

export const getServerSideProps = async (
  _context: GetServerSidePropsContext
): Promise<GetServerSidePropsResult<INoteListPage>> => {
  const MOCK_NOTES: Note[] = [
    {
      id: '1',
      title: 'The Paradigm Mismatch',
      tags: [
        { id: '1', label: 'Spring Boot', color: '#FC7300' },
        { id: '3', label: 'Java', color: '#BFDB38' },
      ],
      status: 'DRAFT',
      emoji: '1f92a',
    },
    {
      id: '2',
      title: 'The Paradigm Mismatch',
      tags: [
        { id: '1', label: 'Spring Boot', color: '#FC7300' },
        { id: '3', label: 'Java', color: '#BFDB38' },
      ],
      status: 'PUBLISHED',
      emoji: '1f92a',
    },
    {
      id: '3',
      title: 'Micronaut Hexa Architecture',
      tags: [
        { id: '4', label: 'Micronaut', color: '#F55050' },
        {
          id: '5',
          label: 'Clean Architecture',
          color: '#FF78F0',
        },
      ],
      status: 'DRAFT',
      emoji: '1f92a',
    },
    {
      id: '4',
      title: 'The Paradigm Mismatch',
      tags: [
        { id: '1', label: 'Spring Boot', color: '#FC7300' },
        { id: '3', label: 'Java', color: '#BFDB38' },
      ],
      status: 'DRAFT',
      emoji: '1f92a',
    },
    {
      id: '5',
      title: 'The Paradigm Mismatch',
      tags: [
        { id: '1', label: 'Spring Boot', color: '#FC7300' },
        { id: '3', label: 'Java', color: '#BFDB38' },
        { id: '3', label: 'Java', color: '#BFDB38' },
      ],
      status: 'PUBLISHED',
      emoji: '1f92a',
    },
    {
      id: '6',
      title: 'Micronaut Hexa Architecture',
      tags: [
        { id: '4', label: 'Micronaut', color: '#F55050' },
        {
          id: '5',
          label: 'Clean Architecture',
          color: '#FF78F0',
        },
      ],
      status: 'DRAFT',
      emoji: '1f92a',
    },
  ];

  const MOCK_COLLECTION_PREVIEW: CollectionPreview[] = [
    {
      id: '1',
      title: 'Java Spring Boot',
      tags: [
        { id: '1', label: 'Array', color: '#FC7300' },
        { id: '2', label: 'Binary Tree', color: '#00425A' },
      ],
    },
    {
      id: '2',
      title: 'Java Spring Boot',
      tags: [
        { id: '1', label: 'Array', color: '#FC7300' },
        { id: '2', label: 'Binary Tree', color: '#00425A' },
      ],
    },
    {
      id: '3',
      title: 'Java Spring Boot',
      tags: [
        { id: '1', label: 'Array', color: '#FC7300' },
        { id: '2', label: 'Binary Tree', color: '#00425A' },
      ],
    },
    {
      id: '4',
      title: 'Java Spring Boot',
      tags: [
        { id: '1', label: 'Array', color: '#FC7300' },
        { id: '2', label: 'Binary Tree', color: '#00425A' },
      ],
    },
    {
      id: '5',
      title: 'Java Spring Boot',
      tags: [
        { id: '1', label: 'Array', color: '#FC7300' },
        { id: '2', label: 'Binary Tree', color: '#00425A' },
      ],
    },
    {
      id: '6',
      title: 'Java Spring Boot',
      tags: [
        { id: '1', label: 'Array', color: '#FC7300' },
        { id: '2', label: 'Binary Tree', color: '#00425A' },
      ],
    },
    {
      id: '7',
      title: 'Java Spring Boot',
      tags: [
        { id: '1', label: 'Array', color: '#FC7300' },
        { id: '2', label: 'Binary Tree', color: '#00425A' },
      ],
    },
    {
      id: '8',
      title: 'Java Spring Boot',
      tags: [
        { id: '1', label: 'Array', color: '#FC7300' },
        { id: '2', label: 'Binary Tree', color: '#00425A' },
      ],
    },
  ];

  return {
    props: {
      notes: MOCK_NOTES,
      collectionPreview: MOCK_COLLECTION_PREVIEW,
    },
  };
};

export default NoteListPage;
