import React, { useState } from "react";
import Modal from "./Component/Modal";
import "./App.css";
import { WalletLinkConnector } from "@web3-react/walletlink-connector";
import { WalletConnectConnector } from "@web3-react/walletconnect-connector";
import { InjectedConnector } from "@web3-react/injected-connector";
import { useWeb3React } from '@web3-react/core'
import CoinbaseIcon from './imgs/Coinbase.png'
import WalletConnectIcon from './imgs/walletconnect.png'
import MetaMaskIcon from './imgs/MetaMask_Fox.png'


export default function App(){
 
  
    
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
  const [toggle, setToggle] = useState(false)
  
    return (
      <div className="App">
        <button
          class="toggle-button"
          id="centered-toggle-button"
          onClick={e => {
            setToggle(true);
       }}
        >
          {" "}
          Connect Wallet{" "}
        </button>
        
        <Modal onClose={e => {
            setToggle(false);
       }} show={toggle}>
        <button style={{padding:5, border:'none'}} onClick={() => { activate(CoinbaseWallet) }}><img style={{width:25, height:25}} src={CoinbaseIcon}></img></button>
        <div></div>
        <button style={{padding:5, border:'none'}} onClick={() => { activate(WalletConnect) }}><img style={{width:25, height:25}} src={WalletConnectIcon}></img></button>
        <div></div>
        <button style={{padding:5, border:'none'}} onClick={() => { activate(Injected) }}><img style={{width:25, height:25}} src={MetaMaskIcon}></img></button>



        </Modal>
      </div>
    );
  }


