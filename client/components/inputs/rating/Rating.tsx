import { Flex, Text } from '@chakra-ui/react';
import { Rating as RatingNumber } from '@lib/types/problem';
import { Rating as MantineRating } from '@mantine/core';

export interface IRating {
  value?: RatingNumber;
  label?: string;
  onChange?: (value: RatingNumber) => void;
  disabled?: boolean;
}

const Rating: React.FC<IRating> = ({
  value = 0,
  onChange = () => {},
  label = '',
  disabled = false,
}) => {
  return (
    <div>
      {label && (
        <Text
          display={'inline-block'}
          lineHeight={'1.55'}
          fontSize={'sm'}
          fontWeight={'500'}
          fontFamily={'heading'}
          color={'gray.200'}
        >
          Rating
        </Text>
      )}
      <Flex height={10} alignItems={'center'}>
        <MantineRating value={value} onChange={onChange} readOnly={disabled} />
      </Flex>
    </div>
  );
};

export default Rating;
