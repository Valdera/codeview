import Sidebar from '@components/sidebar/Sidebar';
import cn from 'classnames';
import styles from './PrimaryLayout.module.scss';

export interface IPrimaryLayout extends React.ComponentPropsWithoutRef<'div'> {
  justify?: 'items-center' | 'items-start';
}

const PrimaryLayout: React.FC<IPrimaryLayout> = ({
  children,
  justify = 'items-center',
  ...divProps
}) => {
  return (
    <>
      <Sidebar>
        <div {...divProps} className={cn(styles.root, `${justify}`)}>
          <main className={'px-5'}>{children}</main>
          <div className={'m-auto'} />
        </div>
      </Sidebar>
    </>
  );
};

export default PrimaryLayout;
