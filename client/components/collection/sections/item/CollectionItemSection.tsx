import { useCollectionStore } from '@lib/stores';
import { DragDropContext, Droppable } from 'react-beautiful-dnd';
import CollectionItemSectionCard from './CollectionItemSectionCard';

export interface ICollectionItemSection {
  disabled?: boolean;
}

const CollectionItemSection: React.FC<ICollectionItemSection> = ({
  disabled = false,
}) => {
  const { collection, reorderCollectionItem } = useCollectionStore();

  if (!collection) return <></>;

  return (
    <DragDropContext
      onDragEnd={({ destination, source }) =>
        reorderCollectionItem(source.index, destination?.index || 0)
      }
    >
      <Droppable droppableId={'dnd-list'} direction={'vertical'}>
        {(provided) => (
          <div
            className={'w-full flex flex-col'}
            {...provided.droppableProps}
            ref={provided.innerRef}
          >
            {collection.items?.map((item, index) => (
              <CollectionItemSectionCard
                key={item.itemId}
                index={index}
                data={{ collectionItem: item }}
                disabled={disabled}
              />
            ))}
            {provided.placeholder}
          </div>
        )}
      </Droppable>
    </DragDropContext>
  );
};

export default CollectionItemSection;
