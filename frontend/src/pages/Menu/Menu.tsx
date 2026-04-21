import React from 'react';
import styles from './Menu.module.css';

interface MenuProps {
  onSelectMode: (mode: 'PVP' | 'PVE') => void;
}

const Menu: React.FC<MenuProps> = ({ onSelectMode }) => {
  return (
    <div className={styles.container}>
      <main className={styles.menuCard}>
        <header className={styles.title}>
          <span className={styles.subtitle}>Welcome to</span>
          <h1 className={styles.titleMain}>
            <span className={styles.tic}>tic.</span>
            <span className={styles.tac}>tac.</span>
            <span className={styles.toe}>toe.</span>
          </h1>
        </header>

        <nav className={styles.buttonGroup}>
          <button 
            className={`${styles.menuButton} ${styles.pvp}`}
            onClick={() => onSelectMode('PVP')}
          >
            PVP Local
            <span className={styles.buttonIcon}>👥</span>
          </button>
          
          <button 
            className={`${styles.menuButton} ${styles.pve}`}
            onClick={() => onSelectMode('PVE')}
          >
            Versus CPU
            <span className={styles.buttonIcon}>🤖</span>
          </button>

          <div className={styles.futureOptions}>
            <button className={`${styles.menuButton} ${styles.disabled}`} disabled>
              Online Battle
              <span>Coming Soon</span>
            </button>
            <button className={`${styles.menuButton} ${styles.disabled}`} disabled>
              Ultimate 9x9
              <span>Locked</span>
            </button>
          </div>
        </nav>
      </main>
    </div>
  );
};

export default Menu;
