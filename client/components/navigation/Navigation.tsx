import { Box, BoxProps, useDisclosure } from '@chakra-ui/react';
import {
  FiCompass,
  FiHome,
  FiSettings,
  FiStar,
  FiTrendingUp,
} from 'react-icons/fi';

import { openSpotlight } from '@mantine/spotlight';
import { ReactNode } from 'react';
import { IconType } from 'react-icons';
import NavigationSidebar from './NavigationSidebar';
import NavigationSpotlightProvider from './NavigationSpotlightProvider';
import NavigationTopbar from './NavigationTopbar';

export interface LinkItem {
  name: string;
  icon: IconType;
  activeUrlPaths: Array<string>;
}

export interface INavigation extends BoxProps {
  linkItems?: Array<LinkItem>;
  children: ReactNode;
}

const Navigation: React.FC<INavigation> = ({
  children,
  linkItems = defaultLinkItems,
  ...rest
}) => {
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <NavigationSpotlightProvider>
      <Box minH={'100vh'} backgroundColor={'foreground'} {...rest}>
        <NavigationSidebar
          isOpen={isOpen}
          onClose={onClose}
          linkItems={linkItems}
        />
        <NavigationTopbar
          onOpen={onOpen}
          onSearchClick={() => openSpotlight()}
        />
        <Box ml={{ base: 0, md: 60 }}>{children}</Box>
      </Box>
    </NavigationSpotlightProvider>
  );
};

export default Navigation;

const defaultLinkItems: Array<LinkItem> = [
  { name: 'Home', icon: FiHome, activeUrlPaths: ['home'] },
  { name: 'Trending', icon: FiTrendingUp, activeUrlPaths: [] },
  { name: 'Explore', icon: FiCompass, activeUrlPaths: [] },
  { name: 'Favourites', icon: FiStar, activeUrlPaths: [] },
  { name: 'Settings', icon: FiSettings, activeUrlPaths: [] },
];
