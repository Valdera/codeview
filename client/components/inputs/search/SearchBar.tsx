import { SearchIcon } from '@chakra-ui/icons';
import {
  Input,
  InputGroup,
  InputLeftElement,
  InputProps,
  InputRightElement,
  Kbd,
} from '@chakra-ui/react';

export interface ISearchBar extends InputProps {}

const SearchBar: React.FC<ISearchBar> = ({ ...props }) => {
  return (
    <InputGroup shadow={'md'} borderRadius={'xl'}>
      <InputLeftElement>
        <SearchIcon color={'gray.500'} />
      </InputLeftElement>
      <InputRightElement>
        <Kbd size={'md'} color={'background'}>
          ↵
        </Kbd>
      </InputRightElement>
      <Input
        borderColor={'transparent'}
        color={'gray.200'}
        type={'text'}
        placeholder={'Search note'}
        focusBorderColor={'transparent'}
        backgroundColor={'background'}
        _placeholder={{ opacity: 1, color: 'none' }}
        _hover={{
          borderColor: 'none',
        }}
        {...props}
      />
    </InputGroup>
  );
};

export default SearchBar;
