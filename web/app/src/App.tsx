import './App.css';
import LoginForm from './components/loginFrom';
import RegisterForm from './components/registerForm';
import { Tabs, TabList, TabPanels, Tab, TabPanel } from '@chakra-ui/react';
function App() {
  return (
    <>
      <Tabs>
        <TabList>
          <Tab>login</Tab>
          <Tab>register</Tab>
        </TabList>

        <TabPanels>
          <TabPanel>
            <LoginForm />
          </TabPanel>
          <TabPanel>
            <RegisterForm />
          </TabPanel>
        </TabPanels>
      </Tabs>
    </>
  );
}

export default App;
