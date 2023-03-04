import { useEffect } from 'react';

function Profile() {
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

export default Profile