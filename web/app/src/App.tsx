import './App.css';
import LoginForm from './components/loginFrom';
import RegisterForm from './components/registerForm';
import {
  Tabs,
  TabList,
  TabPanels,
  Tab,
  TabPanel,
  Flex,
} from '@chakra-ui/react';
function App() {
  return (
    <>
      <Flex justify={'center'} align={'center'}>
        <Tabs w={'fit-content'}>
          <TabList justifyContent={'space-around'}>
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
      </Flex>
    </>
  );
}

export default App;
