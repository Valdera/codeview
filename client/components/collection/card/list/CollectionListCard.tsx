import { HStack } from '@chakra-ui/react';
import LabelTag from '@components/tag/label/LabelTag';
import { Collection } from '@lib/types';
import Image from 'next/image';

export interface ICollectionListCard {
  collection: Collection;
}

const CollectionListCard: React.FC<ICollectionListCard> = ({ collection }) => {
  const date = collection.createdAt
    ? new Date(collection.createdAt)
    : undefined;

  return (
    <article className={'flex bg-white transition hover:shadow-xl'}>
      <div className={'rotate-180 p-2 [writing-mode:_vertical-lr]'}>
        {date && (
          <time
            dateTime={date.toLocaleDateString()}
            className={
              'flex items-center justify-between gap-4 text-xs font-bold uppercase text-gray-900'
            }
          >
            <span className={'text-primary-800'}>{date.getFullYear()}</span>
            <span className={'w-px flex-1 bg-gray-900/10'}></span>
            <span className={'text-primary-800'}>
              {date.toLocaleString('default', { month: 'long' }).slice(0, 3)}{' '}
              {date.getDay()}
            </span>
          </time>
        )}
      </div>

      <div className={'hidden sm:block sm:basis-56'}>
        <Image
          alt={collection.title}
          src={collection.imageUrl}
          width={200}
          height={200}
          className={'aspect-square h-full w-full object-cover'}
        />
      </div>

      <div className={'flex flex-1 flex-col justify-between'}>
        <div
          className={
            'border-l border-gray-900/10 p-4 sm:border-l-transparent sm:p-6'
          }
        >
          <a href={'#'}>
            <h3 className={'font-bold text-lg uppercase text-gray-900'}>
              {collection.title}
            </h3>
          </a>

          <HStack marginY={2}>
            {collection.tags.slice(0, 3).map((t) => (
              <LabelTag key={t.id} tag={t} defaultBg={'primary.800'} />
            ))}
          </HStack>

          <p
            className={
              'mt-2 text-sm leading-relaxed text-gray-700 line-clamp-3'
            }
          >
            {collection.description}
          </p>
        </div>

        <div className={'sm:flex sm:items-end sm:justify-end'}>
          <a
            href={'#'}
            className={
              'block bg-secondary-300 px-5 py-3 text-center text-xs font-bold uppercase text-black transition hover:bg-secondary-400'
            }
          >
            See More
          </a>
        </div>
      </div>
    </article>
  );
};

export default CollectionListCard;
