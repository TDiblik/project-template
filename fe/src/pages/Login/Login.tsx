import {useState} from "react";
import {FaInstagram, FaGithub, FaFacebook, FaGoogle} from "react-icons/fa";
import {config} from "../../api/api";

export default function Login() {
  const [email, setEmail] = useState("");

  console.log(config);

  return (
    <div className="flex min-h-screen items-center justify-center bg-base-200">
      <div className="card w-full max-w-md shadow-xl bg-base-100 p-8">
        <h2 className="text-2xl font-bold text-center mb-6">Login or Sign up</h2>

        {/* Email input */}
        <div className="form-control w-full">
          <label className="label mb-1">
            <span className="label-text">Email</span>
          </label>
          <input
            type="email"
            placeholder="Enter your email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            className="input input-bordered w-full"
          />
        </div>

        {/* Continue button */}
        <button className="btn btn-primary w-full mt-4">Continue</button>

        {/* Divider */}
        <div className="divider">Or continue with</div>

        {/* Social login buttons */}
        <div className="grid grid-cols-2 gap-3">
          <button className="btn btn-outline w-full flex items-center gap-2">
            <FaGoogle /> Google
          </button>
          <button className="btn btn-outline w-full flex items-center gap-2">
            <FaGithub /> GitHub
          </button>
          <button className="btn btn-outline w-full flex items-center gap-2">
            <FaFacebook /> Facebook
          </button>
          <button className="btn btn-outline w-full flex items-center gap-2">
            <FaInstagram /> Instagram
          </button>
        </div>
      </div>
    </div>
  );
}
