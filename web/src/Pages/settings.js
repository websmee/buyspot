import { useEffect } from 'react';

function Settings() {
    useEffect(() => {
        window.bindAll();

        return () => {
            window.unbindAll();
        };
    }, []);

    return (
        <></>
    )
}

export default Settings