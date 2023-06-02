import { useNoteStore } from '@lib/stores';
import { DragDropContext, Droppable } from 'react-beautiful-dnd';
import NoteItemSectionCard from './NoteItemSectionCard';

export interface INoteItemSection {
  disabled?: boolean;
}

const NoteItemSection: React.FC<INoteItemSection> = ({ disabled = false }) => {
  const { reorderNoteItem, note } = useNoteStore();

  if (!note) return <></>;

  return (
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
              <NoteItemSectionCard
                key={item.id}
                index={index}
                data={{ noteItem: item }}
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

export default NoteItemSection;
