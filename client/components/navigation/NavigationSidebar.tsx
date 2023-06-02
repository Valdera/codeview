import {
  Box,
  BoxProps,
  CloseButton,
  Code,
  Drawer,
  DrawerContent,
  Flex,
  Text,
} from '@chakra-ui/react';
import { useRouter } from 'next/router';
import { LinkItem } from './Navigation';
import NavigationSidebarLink from './NavigationSidebarLink';

interface INavigationSidebar {
  isOpen: boolean;
  onClose: () => void;
  linkItems: Array<LinkItem>;
}

const NavigationSidebar: React.FC<INavigationSidebar> = ({
  isOpen,
  onClose,
  linkItems,
}) => {
  return (
    <>
      <NavigationSidebarContent
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
        <DrawerContent color={'white'}>
          <NavigationSidebarContent onClose={onClose} linkItems={linkItems} />
        </DrawerContent>
      </Drawer>
    </>
  );
};

interface INavigationSidebarContent extends BoxProps {
  onClose: () => void;
  linkItems: Array<LinkItem>;
}
const NavigationSidebarContent: React.FC<INavigationSidebarContent> = ({
  onClose,
  linkItems,
  ...rest
}) => {
  const { asPath } = useRouter();
  const pathSubstrings = asPath.split('/');
  const currPath = pathSubstrings[1] != '' ? pathSubstrings[1] : 'home';

  return (
    <Box
      transition={'3s ease'}
      backgroundColor={'foreground'}
      width={{ base: 'full', md: 60 }}
      position={'fixed'}
      height={'full'}
      {...rest}
    >
      <Flex h={'20'} alignItems={'center'} mx={4}>
        <Text
          color={'primary.500'}
          fontSize={'2xl'}
          fontFamily={'monospace'}
          fontWeight={'bold'}
          marginRight={3}
        >
          CodeView
        </Text>
        <Code
          backgroundColor={'background'}
          color={'white'}
          fontWeight={'semibold'}
          paddingX={1}
        >
          1.0.0
        </Code>
        <CloseButton
          marginLeft={'auto'}
          display={{ base: 'flex', md: 'none' }}
          onClick={onClose}
        />
      </Flex>
      {linkItems.map((link) => {
        const isActive = link.activeUrlPaths.indexOf(currPath) != -1;

        return (
          <NavigationSidebarLink
            key={link.name}
            icon={link.icon}
            isActive={isActive}
          >
            {link.name}
          </NavigationSidebarLink>
        );
      })}
    </Box>
  );
};

export default NavigationSidebar;
