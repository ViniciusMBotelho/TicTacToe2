import styles from "./title.module.css";

const Title: React.FC = () => {
  return (
    <div className={styles.mainText}>
      <span className={styles.ticToe}>tic.</span>
      <span className={styles.tac}>tac.</span>
      <span className={styles.ticToe}>toe.</span>
    </div>
  );
};

export default Title;