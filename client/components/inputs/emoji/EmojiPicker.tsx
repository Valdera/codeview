import {
  Flex,
  Modal,
  ModalContent,
  ModalOverlay,
  useDisclosure,
} from '@chakra-ui/react';
import { Emoji, EmojiClickData, EmojiStyle, Theme } from 'emoji-picker-react';
import dynamic from 'next/dynamic';
import React from 'react';

const Picker = dynamic(
  () => {
    return import('emoji-picker-react');
  },
  { ssr: false }
);

export interface IEmojiPicker {
  value: string;
  onChange: ((emoji: EmojiClickData, event: MouseEvent) => void) | undefined;
  disabled?: boolean;
  height?: string;
  width?: string;
}

const EmojiPicker: React.FC<IEmojiPicker> = ({
  value,
  onChange,
  height,
  width,
  disabled = false,
}) => {
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <>
      <Flex
        alignSelf={'end'}
        backgroundColor={'background'}
        justifyContent={'center'}
        alignItems={'center'}
        borderRadius={'sm'}
        width={width ? width : '35px'}
        height={height ? height : '35px'}
        cursor={'pointer'}
        onClick={() => {
          if (!disabled) onOpen();
        }}
        transition={'all .25s'}
        _hover={{
          backgroundColor: `${disabled ? 'background' : 'primary.600'}`,
        }}
      >
        {value && (
          <Emoji unified={value} size={22} emojiStyle={EmojiStyle.TWITTER} />
        )}
      </Flex>

      <Modal
        onClose={onClose}
        isOpen={isOpen}
        isCentered
        motionPreset={'scale'}
      >
        <ModalOverlay />
        <ModalContent
          display={'flex'}
          justifyContent={'center'}
          alignItems={'center'}
          backgroundColor={'transparent'}
          border={'none'}
          shadow={'none'}
          transition={'all .15s'}
        >
          <Picker
            onEmojiClick={(emojiData, evt) => {
              if (onChange) onChange(emojiData, evt);
              onClose();
            }}
            lazyLoadEmojis={true}
            autoFocusSearch={false}
            searchDisabled={true}
            theme={Theme.DARK}
            emojiStyle={EmojiStyle.TWITTER}
          />
        </ModalContent>
      </Modal>
    </>
  );
};

export default EmojiPicker;
