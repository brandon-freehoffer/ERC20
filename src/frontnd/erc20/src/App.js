import {useEffect, useState } from 'react';
import {
  Button,
  useDisclosure,
  HStack,
  VStack,
  Text
} from '@chakra-ui/react';
import { ColorModeSwitcher } from './ColorModeSwitcher';
import WalletModal from './Component/Modal'
import { useWeb3React } from '@web3-react/core'
import { WalletLinkConnector } from "@web3-react/walletlink-connector";
import { WalletConnectConnector } from "@web3-react/walletconnect-connector";
import { InjectedConnector } from "@web3-react/injected-connector";

function App() {

  const CoinbaseWallet = new WalletLinkConnector({
    url: `https://mainnet.infura.io/v3/${process.env.INFURA_KEY}`,
    bridge: "https://bridge.walletconnect.org",
    qrcode: true,
  });
  
  const WalletConnect = new WalletConnectConnector({
    rpcUrl: `https://mainnet.infura.io/v3/${process.env.INFURA_KEY}`,
    bridge: "https://bridge.walletconnect.org",
    qrcode: true,
  });
  
  const Injected = new InjectedConnector({
    supportedChainIds: [1, 3, 4, 5, 42]
  });

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

  const {
    library,
    chainId,
    account,
    activate,
    deactivate,
    active
  } = useWeb3React();
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [verified, setVerified] = useState();
  return (
    <>
    <HStack w="100%" justifyContent="right">
    <ColorModeSwitcher></ColorModeSwitcher>
    </HStack>
    <VStack>
        <HStack w="100%" justifyContent="center">
        {!active ? (
                <Button onClick={onOpen}>Connect Wallet</Button>
              ) : (
                <Button onClick={disconnect}>Disconnect</Button>
              )}
        </HStack>
        <HStack justifyContent="center">
          <Text>{`Account: ${truncateAddress(account)}`}</Text>
        </HStack>
    </VStack>
    <WalletModal isOpen={isOpen} closeModal={onClose}/>
  </>
  );
}

export default App;


