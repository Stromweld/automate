export interface SigninUiSetting {
    isResetPasswordTabVisible: boolean;
    isDisplayNameEditable: boolean;
    isProfileMenu: boolean;
    isTimeformatExist: boolean;
}

class LocalUser implements SigninUiSetting {
  isResetPasswordTabVisible = true;
  isDisplayNameEditable = true;
  isProfileMenu = true;
  isTimeformatExist = true;
}

class SamlUser implements SigninUiSetting {
  isResetPasswordTabVisible = false;
  isDisplayNameEditable = false;
  isProfileMenu = true;
  isTimeformatExist = true;
}

class LdapUser implements SigninUiSetting {
  isResetPasswordTabVisible = false;
  isDisplayNameEditable = false;
  isProfileMenu = true;
  isTimeformatExist = true;
}

export class UISettings {
  local: SigninUiSetting = new LocalUser();
  saml: SigninUiSetting = new SamlUser();
  ldap: SigninUiSetting = new LdapUser();
}
