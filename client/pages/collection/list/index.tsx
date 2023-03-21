import { Grid } from '@chakra-ui/react';
import CollectionListCard from '@components/collection/card/list/CollectionListCard';
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
    <>
      <Grid gridTemplateColumns={'1fr 1fr'} gap={5}>
        {collections.map((c) => (
          <CollectionListCard key={c.id} collection={c} />
        ))}
      </Grid>
    </>
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
          imageUrl:
            'https://images.unsplash.com/photo-1609557927087-f9cf8e88de18?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1740&q=80',
          title: 'Spring Data JPA with Hibernate',
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
          createdAt: '2015-03-25',
        },
        {
          id: '2',
          imageUrl:
            'https://images.unsplash.com/photo-1609557927087-f9cf8e88de18?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1740&q=80',
          title: 'Spring Data JPA with Hibernate',
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
          createdAt: '2015-03-25',
        },
        {
          id: '3',
          imageUrl:
            'https://images.unsplash.com/photo-1609557927087-f9cf8e88de18?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1740&q=80',
          title: 'Spring Data JPA with Hibernate',
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
          createdAt: '2015-03-25',
        },
        {
          id: '4',
          imageUrl:
            'https://images.unsplash.com/photo-1609557927087-f9cf8e88de18?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1740&q=80',
          title: 'Spring Data JPA with Hibernate',
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
          createdAt: '2015-03-25',
        },
      ],
    },
  };
};

export default CollectionListPage;
