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
        <main>
          <div {...divProps} className={cn(styles.root, `${justify}`)}>
            {children}
            <div className={'m-auto'} />
          </div>
        </main>
      </Sidebar>
    </>
  );
};

export default PrimaryLayout;
