import {
  Button,
  Grid,
  GridItem,
  Heading,
  Input,
  InputGroup,
  InputLeftAddon,
  InputRightElement,
} from '@chakra-ui/react';
import RichTextEditor from '@components/rte/RichTextEditor';
import { useNoteStore } from '@lib/stores';
import { NoteItem } from '@lib/types';
import { IconGripVertical } from '@tabler/icons';
import cn from 'classnames';
import { useState } from 'react';
import { Draggable } from 'react-beautiful-dnd';
import s from './NoteItemSection.module.scss';

export interface INoteItemSection {
  index: number;
  data: {
    noteItem: NoteItem;
  };
  disabled?: boolean;
}

const NoteItemSection: React.FC<INoteItemSection> = ({
  index,
  data: { noteItem },
  disabled = false,
}) => {
  const { updateNoteItem } = useNoteStore();

  const [header, setHeader] = useState<string>(noteItem.header);
  const [isEditable, setIsEditable] = useState(false);
  const [isLoading, setIsLoading] = useState(false);

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
          backgroundColor={'white'}
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
            {!disabled && (
              <InputGroup>
                <InputLeftAddon>Heading</InputLeftAddon>
                <Input
                  type={'text'}
                  placeholder={'Heading'}
                  value={header}
                  onChange={(evt) => setHeader(evt.target.value)}
                  readOnly={!isEditable}
                />
                <InputRightElement width={'4.5rem'}>
                  {!isEditable ? (
                    <Button
                      h={'1.75rem'}
                      size={'sm'}
                      isLoading={isLoading}
                      onClick={() => {
                        setIsEditable(true);
                      }}
                    >
                      Edit
                    </Button>
                  ) : (
                    <Button
                      h={'1.75rem'}
                      size={'sm'}
                      isLoading={isLoading}
                      onClick={() => {
                        setIsLoading(true);
                        updateNoteItem(noteItem.id, {
                          header,
                          content: noteItem.content,
                        });
                        setIsLoading(false);
                        setIsEditable(false);
                      }}
                    >
                      Save
                    </Button>
                  )}
                </InputRightElement>
              </InputGroup>
            )}
            <Heading size={'md'} color={'primary.700'}>
              {noteItem.header}
            </Heading>
            <RichTextEditor
              content={noteItem.content}
              onSave={(content) => {
                updateNoteItem(noteItem.id, {
                  header: noteItem.header,
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

export default NoteItemSection;
