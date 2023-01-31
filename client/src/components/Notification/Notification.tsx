
import { ToastContainer, ToastPosition } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
export interface NotificationInterface {
	position: string;
}

const Notification : React.FC<NotificationInterface> = ({position}) => {
	return <ToastContainer position={position as ToastPosition} autoClose={2000} hideProgressBar={true}  newestOnTop={false}
	closeOnClick theme="light"/>
};

export default Notification;
