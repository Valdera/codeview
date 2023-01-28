import {
  Box,
  BoxProps,
  CloseButton,
  Drawer,
  DrawerContent,
  Flex,
  Text,
  useColorModeValue,
  useDisclosure,
} from '@chakra-ui/react';
import {
  FiCompass,
  FiHome,
  FiSettings,
  FiStar,
  FiTrendingUp,
} from 'react-icons/fi';

import { ReactNode } from 'react';
import { IconType } from 'react-icons';
import MobileNav from './MobileNav';
import NavItem from './NavItem';

const defaultLinkItems: LinkItems = [
  { name: 'Home', icon: FiHome },
  { name: 'Trending', icon: FiTrendingUp },
  { name: 'Explore', icon: FiCompass },
  { name: 'Favourites', icon: FiStar },
  { name: 'Settings', icon: FiSettings },
];

export interface LinkItem {
  name: string;
  icon: IconType;
}

export type LinkItems = Array<LinkItem>;

export interface ISidebar extends BoxProps {
  linkItems?: LinkItems;
  children: ReactNode;
}

const Sidebar: React.FC<ISidebar> = ({
  children,
  linkItems = defaultLinkItems,
  ...rest
}) => {
  const { isOpen, onOpen, onClose } = useDisclosure();
  return (
    <Box
      minH={'100vh'}
      bg={useColorModeValue('gray.100', 'gray.900')}
      {...rest}
    >
      <SidebarContent
        onClose={() => onClose}
        linkItems={linkItems}
        display={{ base: 'none', md: 'block' }}
      />
      <Drawer
        autoFocus={false}
        isOpen={isOpen}
        placement={'left'}
        onClose={onClose}
        returnFocusOnClose={false}
        onOverlayClick={onClose}
        size={'full'}
      >
        <DrawerContent>
          <SidebarContent onClose={onClose} linkItems={linkItems} />
        </DrawerContent>
      </Drawer>

      {/* mobilenav */}
      <MobileNav onOpen={onOpen} />
      <Box ml={{ base: 0, md: 60 }} p={'4'}>
        {children}
      </Box>
    </Box>
  );
};

interface ISidebarContent extends BoxProps {
  onClose: () => void;
  linkItems: LinkItems;
}

const SidebarContent: React.FC<ISidebarContent> = ({
  onClose,
  linkItems,
  ...rest
}) => {
  return (
    <Box
      transition={'3s ease'}
      bg={useColorModeValue('white', 'gray.900')}
      borderRight={'1px'}
      borderRightColor={useColorModeValue('gray.200', 'gray.700')}
      w={{ base: 'full', md: 60 }}
      pos={'fixed'}
      h={'full'}
      {...rest}
    >
      <Flex
        h={'20'}
        alignItems={'center'}
        mx={'8'}
        justifyContent={'space-between'}
      >
        <Text fontSize={'2xl'} fontFamily={'monospace'} fontWeight={'bold'}>
          Logo
        </Text>
        <CloseButton display={{ base: 'flex', md: 'none' }} onClick={onClose} />
      </Flex>
      {linkItems.map((link) => (
        <NavItem key={link.name} icon={link.icon}>
          {link.name}
        </NavItem>
      ))}
    </Box>
  );
};

export default Sidebar;
