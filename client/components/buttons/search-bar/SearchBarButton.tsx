import { SearchIcon } from '@chakra-ui/icons';
import {
  Input,
  InputGroup,
  InputLeftElement,
  InputRightElement,
  Kbd,
} from '@chakra-ui/react';

export interface ISearchBarButton {
  placeholder?: string;
  onClick?: () => void;
}
const SearchBarButton: React.FC<ISearchBarButton> = ({
  onClick = () => {},
  placeholder = '',
}) => {
  return (
    <InputGroup shadow={'md'} borderRadius={'xl'}>
      <InputLeftElement>
        <SearchIcon color={'gray.500'} />
      </InputLeftElement>
      <InputRightElement
        display={{ base: 'none', sm: 'flex' }}
        width={'7rem'}
        color={'white'}
        gap={1}
      >
        <Kbd color={'background'}>shift</Kbd> +{' '}
        <Kbd color={'background'}>H</Kbd>
      </InputRightElement>
      <Input
        cursor={'pointer'}
        type={'text'}
        placeholder={placeholder}
        onClick={onClick}
        focusBorderColor={'transparent'}
        backgroundColor={'background'}
        borderColor={'none'}
        _placeholder={{ opacity: 1, color: 'none' }}
        _hover={{
          borderColor: 'none',
        }}
        isReadOnly
      />
    </InputGroup>
  );
};

export default SearchBarButton;
