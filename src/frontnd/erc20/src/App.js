import {useEffect, useState } from 'react';
import {
  Button,
  useDisclosure,
  HStack,
  VStack,
  Text,
  NumberInput,
  NumberInputField,
  NumberInputStepper,
  NumberDecrementStepper,
  NumberIncrementStepper,
  interactivity
} from '@chakra-ui/react';
import { ColorModeSwitcher } from './ColorModeSwitcher';
import WalletModal from './Component/Modal'
import { useWeb3React } from '@web3-react/core'




function App() {
  const [appState, setAppState] = useState({
    loading: false,
    tokenInfo: null,
    textChanged: false
  });
 
  const handleInput = async (e) => {
   
    console.log(e.target.value);
    const foo = await fetch('http://127.0.0.1:1351/Sign?address=' + account + '&amount=' + e.target.value)
    .then((response) => response.json())
    const msg = JSON.stringify(foo);
    setMessage(e.target.value);
  };


  
  const [signature, setSignature] = useState("");
  useEffect(() => {
     setAppState({ loading: true });
     fetch('http://127.0.0.1:1351/GetTokenInfo')
     .then((response) => response.json())
     .then((data) => { setAppState({ loading: false, tokenInfo: data });
    });
  }, [setAppState]);

  const refreshState = () => {
    window.localStorage.setItem("provider", undefined);
    
    setVerified(undefined);
  };
  const disconnect = () => {
    refreshState();
    deactivate();
  };
  const truncateAddress = (address) => {
  if (!address) return "No Account";
  const match = address.match(
    /^(0x[a-zA-Z0-9]{2})[a-zA-Z0-9]+([a-zA-Z0-9]{2})$/
  );
  if (!match) return address;
  return `${match[1]}…${match[2]}`;
};
const signMessage = async () => {
 
  if (!library) return;
  try {
    const signature = await library.provider.request({
      method: "personal_sign",
      params: [message, account]
    });
    const parse = parseInt(message, 10)
    if(!isNaN(parse))
    {
      const foo = await fetch('http://127.0.0.1:1351/Transfer?address=' + account + "&amount=" + message)
      .then((response) => response.json())
      setSignedMessage(message);
      setSignature(signature);
    }
  } catch (error) {
    setError(error);
  }
};


const {
  library,
  chainId,
  account,
  activate,
  deactivate,
  active
} = useWeb3React();
  const [signedMessage, setSignedMessage] = useState("");
  const [message, setMessage] = useState("");
  const [error, setError] = useState("");
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [verified, setVerified] = useState();
  if(appState.loading)
  {
    <Text>Loading..</Text>
  }
  if(appState.tokenInfo != null)
  {
  return (
    <>
    <HStack w="100%" justifyContent="right">
    <ColorModeSwitcher></ColorModeSwitcher>
    </HStack>
    
    <VStack>
        <HStack w="100%" justifyContent="center">
            <Text>
            Token Name: {appState.tokenInfo.Name}
            </Text>
        </HStack>
        <HStack w="100%" justifyContent="center">
            <Text>
            Symbol: {appState.tokenInfo.Symbol}
            </Text>
        </HStack>
        <HStack w="100%" justifyContent="center">
        {!active ? (
                <Button colorScheme="teal" variant="outline" onClick={onOpen}>Connect Wallet</Button>
              ) : (
                <Button colorScheme="red" variant="outline" onClick={disconnect}>Disconnect</Button>
              )}
        </HStack>
        <HStack>
    
        </HStack>
        <HStack justifyContent="center">
          <Text>{`Account: ${truncateAddress(account)}`}</Text>
        </HStack>
        {active ? (
        <HStack w="50%" justifyContent="center">
                <Text w="25%">Get Tokens:</Text>
                <NumberInput
                  value={appState.amtVal}
                  w="50%"
                  defaultValue={0}
                  max={1000000000}
                  keepWithinRange={false}
                  clampValueOnBlur={false}
                  >
                    <NumberInputField onChange={handleInput}/>
                    <NumberInputStepper>
                      <NumberIncrementStepper />
                      <NumberDecrementStepper />
                    </NumberInputStepper>
                  </NumberInput>
            
                <Button w="25%" onClick={signMessage}>Mint Tokens</Button>
                 
        </HStack> ) : (<HStack></HStack>)
        }   
    </VStack>
    <WalletModal isOpen={isOpen} closeModal={onClose}/>
    
  </>
  );
}
}


export default App;


