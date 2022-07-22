import React from "react";
import { WalletLinkConnector } from "@web3-react/walletlink-connector";
import { WalletConnectConnector } from "@web3-react/walletconnect-connector";
import { InjectedConnector } from "@web3-react/injected-connector";
import { useWeb3React } from '@web3-react/core'
import CoinbaseIcon from '../../imgs/Coinbase.png'
import WalletConnectIcon from '../../imgs/walletconnect.png'
import MetaMaskIcon from '../../imgs/MetaMask_Fox.png'
import { 
  VStack,
  HStack,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalBody,
  ModalCloseButton,
  Button,
  Text,
  Image,
  useDisclosure,
 } from "@chakra-ui/react";

export default function WalletModal({ isOpen, closeModal }){
  const CoinbaseWallet = new WalletLinkConnector({
    rpcUrl: `https://mainnet.infura.io/v3/${process.env.INFURA_KEY}`,
    bridge: "https://bridge.walletconnect.org",
    qrcode: true
  });
  
  const WalletConnect = new WalletConnectConnector({
   rpcUrl: `https://mainnet.infura.io/v3/${process.env.INFURA_KEY}`,
   bridge: "https://bridge.walletconnect.org",
   qrcode: true,
  });
  
  const Injected = new InjectedConnector({
   supportedChainIds: [1, 3, 4, 5, 42, 1337]
  }); 

  
  const { activate, deactivate } = useWeb3React();
  const { onClose } = useDisclosure()
   
    return (
      <Modal isOpen={isOpen} onClose={closeModal} isCentered>
      <ModalOverlay />
      <ModalContent w="300px">
        <ModalHeader>Select Wallet</ModalHeader>
        <ModalCloseButton
          _focus={{
            boxShadow: "none"
          }}
        />
        <ModalBody paddingBottom="1.5rem">
          <VStack>
            <Button
              variant="outline"
              onClick={() => {
                activate(CoinbaseWallet);
        
                closeModal();
              }}
              w="100%"
            >
              <HStack w="100%" justifyContent="center">
                <Image
                  src={CoinbaseIcon}
                  alt="Coinbase Wallet Logo"
                  width={25}
                  height={25}
                  borderRadius="3px"
                />
                <Text>Coinbase Wallet</Text>
              </HStack>
            </Button>
            <Button
              variant="outline"
              onClick={() => {
                activate(WalletConnect);
              
                closeModal();
              }}
              w="100%"
            >
              <HStack w="100%" justifyContent="center">
                <Image
                  src={WalletConnectIcon}
                  alt="Wallet Connect Logo"
                  width={26}
                  height={26}
                  borderRadius="3px"
                />
                <Text>Wallet Connect</Text>
              </HStack>
            </Button>
            <Button
              variant="outline"
              onClick={() => {
                  activate(Injected);
           
                closeModal();
              }}
              w="100%"
            >
              <HStack w="100%" justifyContent="center">
                <Image
                  src={MetaMaskIcon}
                  alt="Metamask Logo"
                  width={25}
                  height={25}
                  borderRadius="3px"
                />
                <Text>Metamask</Text>
              </HStack>
            </Button>
          </VStack>
        </ModalBody>
      </ModalContent>
    </Modal>
    );
  }