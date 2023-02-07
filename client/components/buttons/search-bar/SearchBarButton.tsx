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
      <InputRightElement width={'7rem'}>
        <Kbd>shift</Kbd> + <Kbd>H</Kbd>
      </InputRightElement>
      <Input
        cursor={'pointer'}
        isReadOnly
        type={'text'}
        placeholder={placeholder}
        onClick={onClick}
      />
    </InputGroup>
  );
};

export default SearchBarButton;
