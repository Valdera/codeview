import { Grid, VStack } from '@chakra-ui/react';
import CollectionListCard from '@components/collection/cards/list/CollectionListCard';
import CollectionFilterSection from '@components/collection/sections/filter/CollectionFilterSection';
import PrimaryLayout from '@components/layouts/primary/PrimaryLayout';
import { Collection } from '@lib/types';
import { NextPageWithLayout } from '@pages/page';
import { GetServerSidePropsContext, GetServerSidePropsResult } from 'next';

interface ICollectionListPage {
  collections: Collection[];
}

const CollectionListPage: NextPageWithLayout<ICollectionListPage> = ({
  collections,
}) => {
  return (
    <VStack width={'full'} gap={5}>
      <CollectionFilterSection />
      <Grid
        gridTemplateColumns={{ base: '1fr', lg: '1fr 1fr', xl: '1fr 1fr 1fr' }}
        gap={5}
      >
        {collections.map((c) => (
          <CollectionListCard key={c.id} collection={c} />
        ))}
      </Grid>
    </VStack>
  );
};

CollectionListPage.getLayout = (page) => {
  return <PrimaryLayout justify={'items-start'}>{page}</PrimaryLayout>;
};

export const getServerSideProps = async (
  _context: GetServerSidePropsContext
): Promise<GetServerSidePropsResult<ICollectionListPage>> => {
  return {
    props: {
      collections: [
        {
          id: '1',
          title: 'Spring Data JPA with Hibernate',
          emoji: '1f4d1',
          type: 'PROBLEM',
          description:
            'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Recusandae dolores, possimus pariatur animi temporibus nesciunt praesentium dolore sed nulla ipsum eveniet corporis quidem, mollitia itaque minus soluta, voluptates neque explicabo tempora nisi culpa eiusatque dignissimos. Molestias explicabo corporis voluptatem?',
          tags: [
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
          ],
        },
        {
          id: '2',
          title: 'Spring Data JPA with Hibernate',
          emoji: '1f4d3',
          type: 'NOTE',
          description:
            'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Recusandae dolores, possimus pariatur animi temporibus nesciunt praesentium dolore sed nulla ipsum eveniet corporis quidem, mollitia itaque minus soluta, voluptates neque explicabo tempora nisi culpa eiusatque dignissimos. Molestias explicabo corporis voluptatem?',
          tags: [
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
          ],
        },
        {
          id: '3',
          title: 'Spring Data JPA with Hibernate',
          emoji: '1f4d3',
          type: 'NOTE',
          description:
            'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Recusandae dolores, possimus pariatur animi temporibus nesciunt praesentium dolore sed nulla ipsum eveniet corporis quidem, mollitia itaque minus soluta, voluptates neque explicabo tempora nisi culpa eiusatque dignissimos. Molestias explicabo corporis voluptatem?',
          tags: [
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
          ],
        },
        {
          id: '4',
          title: 'Spring Data JPA with Hibernate',
          emoji: '1f4d1',
          type: 'PROBLEM',
          description:
            'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Recusandae dolores, possimus pariatur animi temporibus nesciunt praesentium dolore sed nulla ipsum eveniet corporis quidem, mollitia itaque minus soluta, voluptates neque explicabo tempora nisi culpa eiusatque dignissimos. Molestias explicabo corporis voluptatem?',
          tags: [
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
          ],
        },
      ],
    },
  };
};

export default CollectionListPage;
