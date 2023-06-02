import {
  Box,
  FlexProps,
  Grid,
  GridItem,
  HStack,
  IconButton,
} from '@chakra-ui/react';
import SearchBarButton from '@components/buttons/search-bar/SearchBarButton';
import { IconBrandDiscord, IconBrandGithub } from '@tabler/icons';
import { FiMenu } from 'react-icons/fi';

interface INavigationTopbar extends FlexProps {
  onOpen: () => void;
  onSearchClick: () => void;
}
const NavigationTopbar: React.FC<INavigationTopbar> = ({
  onOpen,
  onSearchClick,
  ...rest
}) => {
  return (
    <Grid
      marginLeft={{ base: 0, md: 60 }}
      paddingX={{ base: 4, md: 4 }}
      height={20}
      alignItems={'center'}
      templateColumns={{ base: 'max-content 1fr', md: '1fr max-content' }}
      backgroundColor={'foreground'}
      justifyContent={{ base: 'space-between', md: 'flex-end' }}
      {...rest}
    >
      <IconButton
        display={{ base: 'flex', md: 'none' }}
        onClick={onOpen}
        variant={'solid'}
        colorScheme={'blackAlpha'}
        aria-label={'open menu'}
        icon={<FiMenu />}
      />
      <GridItem
        display={'flex'}
        justifyContent={'center'}
        alignItems={'center'}
      >
        <Box width={{ base: '80%', md: '70%', lg: '60%' }}>
          <SearchBarButton
            placeholder={'Search or jump to...'}
            onClick={onSearchClick}
          />
        </Box>
      </GridItem>
      <GridItem display={{ base: 'none', md: 'block' }}>
        <HStack>
          <IconButton
            size={'lg'}
            variant={'solid'}
            colorScheme={'blackAlpha'}
            aria-label={'Github Button'}
            icon={<IconBrandGithub />}
          />
          <IconButton
            size={'lg'}
            variant={'solid'}
            colorScheme={'purple'}
            aria-label={'Discord Button'}
            icon={<IconBrandDiscord />}
          />
        </HStack>
      </GridItem>
    </Grid>
  );
};

export default NavigationTopbar;
