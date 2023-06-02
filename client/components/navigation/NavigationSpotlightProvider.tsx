import { SpotlightAction, SpotlightProvider } from '@mantine/spotlight';
import {
  IconDashboard,
  IconFileText,
  IconHome,
  IconSearch,
} from '@tabler/icons';
import { ReactNode } from 'react';

export interface INavigationSpotlightProvider {
  children: ReactNode;
}

const NavigationSpotlightProvider: React.FC<INavigationSpotlightProvider> = ({
  children,
}) => {
  return (
    <SpotlightProvider
      actions={actions}
      searchIcon={<IconSearch size={18} />}
      searchPlaceholder={'Search...'}
      shortcut={'mod + shift + 1'}
      nothingFoundMessage={'Nothing found...'}
    >
      {children}
    </SpotlightProvider>
  );
};

export default NavigationSpotlightProvider;

const actions: SpotlightAction[] = [
  {
    title: 'Home',
    description: 'Get to home page',
    onTrigger: () => console.log('Home'),
    icon: <IconHome size={18} />,
  },
  {
    title: 'Dashboard',
    description: 'Get full information about current system status',
    onTrigger: () => console.log('Dashboard'),
    icon: <IconDashboard size={18} />,
  },
  {
    title: 'Documentation',
    description: 'Visit documentation to lean more about all features',
    onTrigger: () => console.log('Documentation'),
    icon: <IconFileText size={18} />,
  },
];
