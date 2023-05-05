import OneSignal from 'react-onesignal';

export default async function runOneSignal() {
    await OneSignal.init({ appId: '7447ecbb-a4df-4023-88a6-db2efdc5bdc4', allowLocalhostAsSecureOrigin: true});
    OneSignal.showSlidedownPrompt();
}