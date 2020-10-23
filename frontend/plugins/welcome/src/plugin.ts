import { createPlugin } from '@backstage/core';
import Login  from './components/Login';
import WelcomePage from './components/WelcomePage';
import CreateUser from './components/Users';
import Tables from './components/Table';
export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', Login);
    router.registerRoute('/WelcomePage', WelcomePage);
    router.registerRoute('/user', CreateUser);
    router.registerRoute('/table', Tables);
  },
});