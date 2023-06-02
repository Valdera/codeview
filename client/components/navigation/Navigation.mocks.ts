import {
  FiCompass,
  FiHome,
  FiSettings,
  FiStar,
  FiTrendingUp,
} from 'react-icons/fi';
import { INavigation } from './Navigation';

const base: INavigation = {
  children: '{{component}}',
  linkItems: [
    { name: 'Home', icon: FiHome, activeUrlPaths: ['home'] },
    { name: 'Trending', icon: FiTrendingUp, activeUrlPaths: [] },
    { name: 'Explore', icon: FiCompass, activeUrlPaths: [] },
    { name: 'Favourites', icon: FiStar, activeUrlPaths: [] },
    { name: 'Settings', icon: FiSettings, activeUrlPaths: [] },
  ],
};

export const mockNavigationProps = {
  base,
};
