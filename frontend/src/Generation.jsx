import React, { useState } from 'react';
import './PasswordGenerationPage.css'; 
import { GeneratePassword } from '../wailsjs/go/app/app';
import { faCopy } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

function PasswordGenerationPage() {
  const [formData, setFormData] = useState({
    password: '',
    domain: '',
    prompts: '',
    length: '16',
  });

  const [showPassword, setShowPassword] = useState(false);
  const [generatedPassword, setGeneratedPassword] = useState(''); 
  const [showResult, setShowResult] = useState(false); 
  const [copyMessageVisible, setCopyMessageVisible] = useState(false);
  const [errors, setErrors] = useState({
    domain: '',
  });

  function validateDomain(domain) {
    const allowedCharacters = /^[A-Za-z0-9\-.]+$/;
    const validTLDs = [
      'com', 'net', 'org', 'info', 'io', 'app', 'blog', 'tech', 'co', 'edu',
      'gov', 'mil', 'int', 'biz', 'jobs', 'mobi', 'name', 'ly', 'tel', 'tv',
      'asia', 'cat', 'pro', 'aero', 'coop', 'museum', 'arpa', 'travel', 'xxx',
    ];
  
    if (!allowedCharacters.test(domain)) {
      return 'Domain contains invalid characters';
    }
  
    const domainParts = domain.split('.');
    if (domainParts.length < 2) {
      return 'Domain must have at least one dot';
    }
    const tld = domainParts[domainParts.length - 1];
    if (tld.length < 2 || tld.length > 6 || !validTLDs.includes(tld)) {
      return 'Invalid Top-Level Domain (TLD)';
    }
  
    return '';
  }

  const handleChange = (e) => {
    const { name, value } = e.target;
    const errorMessage = name === 'domain' ? validateDomain(value) : '';

    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));

    setErrors((prevErrors) => ({
      ...prevErrors,
      [name]: errorMessage,
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    
    if (errors.domain){
      return
    }

    let promts = formData.prompts.split(" ")
    GeneratePassword(formData.password, formData.domain, formData.length, promts).then((result)=> {
      setGeneratedPassword(result);
      setShowResult(true); 
    });
  };

  const handleCopyToClipboard = () => {
    navigator.clipboard.writeText(generatedPassword);
    setCopyMessageVisible(true);

    setTimeout(() => {
      setCopyMessageVisible(false);
    }, 1000);
  };


  function extractDomainFromUrl(url) {
    try {
      const parsedUrl = new URL(url);
      return parsedUrl.hostname;
    } catch (error) {
      return 'Invalid URL:' + error;
    }
  }

  const handlePaste = (e) => {
    e.preventDefault();
    const pastedText = e.clipboardData.getData('text');
    const extractDomain = extractDomainFromUrl(pastedText);
    if (extractDomain) {
      setFormData((prevData) => ({
        ...prevData,
        domain: extractDomain,
      }));
    } 
  };

  return (
    <div className="form-page">
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="password">Master Password</label>
          <div className="password-input">
              <input
                type={showPassword ? 'text' : 'password'}
                id="password"
                name="password"
                value={formData.password}
                onChange={handleChange}
                required
              />
              <button
                type="button"
                onClick={() => setShowPassword(!showPassword)}
                className="show-password-button"
              >
                Show
              </button>
          </div>
        </div>
        <div className="form-group">
          <label htmlFor="domain">Site Domain</label>
          <input
            type="text"
            id="domain"
            name="domain"
            value={formData.domain}
            onChange={handleChange}
            onPaste={handlePaste}
            className={errors.domain ? 'error' : ''}
            required
          />
          {errors.domain && <span className="error-message">{errors.domain}</span>}
        </div>
        <div className="form-group">
          <label htmlFor="prompts">Prompts</label>
          <input
            type="text"
            id="prompts"
            name="prompts"
            value={formData.prompts}
            onChange={handleChange}
          />
        </div>
        <div className="form-group">
          <label>Length</label>
          <div className="radio-group">
            <label>
              <input
                type="radio"
                name="length"
                value="16"
                checked={formData.length === '16'}
                onChange={handleChange}
              />
              16
            </label>
            <label>
              <input
                type="radio"
                name="length"
                value="32"
                checked={formData.length === '32'}
                onChange={handleChange}
              />
              32
            </label>
          </div>
        </div>
        <button type="submit" className="submit-button">
          Generate
        </button>
      </form>

      {showResult && (
        <div className="result-container">
          <p>Your Password: <b>{generatedPassword}</b></p>
          <button
            type="button"
            className="copy-to-clipboard-button"
            onClick={handleCopyToClipboard}
          >
            {copyMessageVisible ? 'Copy!' : <FontAwesomeIcon icon={faCopy} />}
          </button>
        </div>
      )}

    </div>
  );
}

export default PasswordGenerationPage;