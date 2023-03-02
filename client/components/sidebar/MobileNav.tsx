import {
  Box,
  FlexProps,
  Grid,
  GridItem,
  HStack,
  IconButton,
  useColorModeValue,
} from '@chakra-ui/react';
import SearchBarButton from '@components/buttons/search-bar/SearchBarButton';
import { IconBrandDiscord, IconBrandGithub } from '@tabler/icons';
import { FiMenu } from 'react-icons/fi';

interface MobileProps extends FlexProps {
  onOpen: () => void;
  onSearchClick: () => void;
}
const MobileNav = ({ onOpen, onSearchClick, ...rest }: MobileProps) => {
  return (
    <Grid
      ml={{ base: 0, md: 60 }}
      px={{ base: 4, md: 4 }}
      height={'20'}
      alignItems={'center'}
      templateColumns={{ base: 'max-content 1fr', md: '1fr max-content' }}
      bg={useColorModeValue('white', 'gray.900')}
      borderBottomWidth={'1px'}
      borderBottomColor={useColorModeValue('primary.100', 'gray.700')}
      justifyContent={{ base: 'space-between', md: 'flex-end' }}
      {...rest}
    >
      <IconButton
        display={{ base: 'flex', md: 'none' }}
        onClick={onOpen}
        variant={'outline'}
        aria-label={'open menu'}
        icon={<FiMenu />}
      />
      <GridItem
        display={'flex'}
        justifyContent={'center'}
        alignItems={'center'}
      >
        <Box width={{ base: '80%', md: '60%', lg: '50%' }}>
          <SearchBarButton
            placeholder={'Search problem'}
            onClick={onSearchClick}
          />
        </Box>
      </GridItem>
      <GridItem display={{ base: 'none', md: 'block' }}>
        <HStack>
          <IconButton
            variant={'outline'}
            colorScheme={'blackAlpha'}
            aria-label={'Github Button'}
            icon={<IconBrandGithub />}
          />
          <IconButton
            variant={'outline'}
            colorScheme={'purple'}
            aria-label={'Discord Button'}
            icon={<IconBrandDiscord />}
          />
        </HStack>
      </GridItem>
    </Grid>
  );
};

export default MobileNav;
