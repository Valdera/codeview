import { Flex, HStack, IconButton } from '@chakra-ui/react';
import { useAppStore } from '@lib/store';
import { GrAdd, GrSubtract } from 'react-icons/gr';
import s from './Counter.module.scss';

const Counter: React.FC = () => {
  const { counter, increment, decrement } = useAppStore();

  return (
    <>
      <Flex direction={'column'} alignItems={'center'} gap={'5'}>
        <div className={s.display}>{counter.current}</div>
        <HStack gap={5}>
          <IconButton
            colorScheme={'whiteAlpha'}
            shadow={'lg'}
            aria-label={'decrement counter'}
            icon={<GrSubtract />}
            _hover={{
              backgroundColor: 'teal.100',
            }}
            onClick={decrement}
          />
          <IconButton
            colorScheme={'whiteAlpha'}
            shadow={'lg'}
            aria-label={'increment counter'}
            icon={<GrAdd />}
            _hover={{
              backgroundColor: 'teal.100',
            }}
            onClick={increment}
          />
        </HStack>
      </Flex>
    </>
  );
};

export default Counter;
