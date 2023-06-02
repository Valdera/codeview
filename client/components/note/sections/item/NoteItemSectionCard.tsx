import { Grid, GridItem } from '@chakra-ui/react';
import RichTextEditor from '@components/inputs/rte/RichTextEditor';
import { useNoteStore } from '@lib/stores';
import { NoteItem } from '@lib/types';
import { IconGripVertical } from '@tabler/icons';
import cn from 'classnames';
import { Draggable } from 'react-beautiful-dnd';
import s from './NoteItemSectionCard.module.scss';

export interface INoteItemSectionCard {
  index: number;
  data: {
    noteItem: NoteItem;
  };
  disabled?: boolean;
}

const NoteItemSectionCard: React.FC<INoteItemSectionCard> = ({
  index,
  data: { noteItem },
  disabled = false,
}) => {
  const { updateNoteItem } = useNoteStore();

  return (
    <Draggable
      isDragDisabled={disabled}
      key={noteItem.id}
      index={index}
      draggableId={noteItem.id}
    >
      {(provided, snapshot) => (
        <Grid
          gap={2}
          width={'full'}
          rounded={'md'}
          gridTemplateColumns={'max-content 1fr'}
          marginBottom={3}
          backgroundColor={'foreground'}
          className={cn(s.item, {
            [s.itemDragging]: snapshot.isDragging,
          })}
          ref={provided.innerRef}
          {...provided.draggableProps}
        >
          {!disabled && (
            <GridItem className={s.dragHandle} {...provided.dragHandleProps}>
              <IconGripVertical color={'white'} size={18} stroke={1.5} />
            </GridItem>
          )}
          <GridItem
            width={'full'}
            display={'flex'}
            flexDir={'column'}
            gap={5}
            padding={'5'}
            colSpan={disabled ? 2 : 1}
          >
            <RichTextEditor
              content={noteItem.content}
              onSave={(content) => {
                updateNoteItem(noteItem.id, {
                  content: content,
                });
              }}
              disabled={disabled}
            />
          </GridItem>
        </Grid>
      )}
    </Draggable>
  );
};

export default NoteItemSectionCard;
