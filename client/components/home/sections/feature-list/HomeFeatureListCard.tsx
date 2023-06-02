import {
  Box,
  Button,
  Flex,
  Heading,
  Text,
  useColorModeValue,
  VStack,
} from '@chakra-ui/react';
import Link from 'next/link';
import { ReactElement } from 'react';

export interface IHomeFeatureListCard {
  heading: string;
  description: string;
  icon: ReactElement;
  href: string;
}

export const HomeFeatureListCard: React.FC<IHomeFeatureListCard> = ({
  heading,
  description,
  icon,
  href,
}) => {
  return (
    <VStack
      align={'start'}
      width={'full'}
      borderWidth={'1px'}
      borderRadius={'lg'}
      backgroundColor={'white'}
      overflow={'hidden'}
      padding={5}
    >
      <Flex
        width={16}
        height={16}
        align={'center'}
        justify={'center'}
        color={'white'}
        rounded={'full'}
        marginBottom={3}
        bg={useColorModeValue('gray.100', 'gray.700')}
      >
        {icon}
      </Flex>
      <Box marginTop={2} paddingBottom={2}>
        <Heading size={'md'} color={'primary.700'}>
          {heading}
        </Heading>
        <Text marginTop={1} fontSize={'sm'}>
          {description}
        </Text>
      </Box>
      <Link href={href} className={'!mt-auto'}>
        <Button variant={'link'} colorScheme={'secondary'} size={'sm'}>
          Learn more
        </Button>
      </Link>
    </VStack>
  );
};

export default HomeFeatureListCard;
