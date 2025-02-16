/* eslint-disable @typescript-eslint/ban-ts-comment */
type JWTPayload = {
  region_id: string;
  aud?: string;
  exp?: number;
  iat?: number;
  iss?: string;
  roles?: Array<string>;
  sub?: string;
};

function isTemproary(token: string) {
  const payload = getPayload(token);
  // @ts-ignore
  if (payload.iss.startsWith("social")) {
    return true;
  }

  return false;
}

function isValid(token: string) {
  const payload = getPayload(token);
  if (payload) {
    return true;
  } else {
    return false;
  }
}

function isExpired(token: string) {
  const payload = getPayload(token);
  if (payload) {

    //@ts-ignore
    const expDate = new Date(payload.exp * 1000)
    const currentDate = new Date()
  
    if (expDate > currentDate) {
      return false;
    }

    return true;
  } else {
    return true;
  }

}

function getPayload(token: string) {
  try {
    const payloadPart = token.split(".")[1];
    const payload: JWTPayload = JSON.parse(atob(payloadPart));
    return payload;
    //@ts-ignore
  } catch (e: unknwn) {
    console.log(e)
    return {};
  }
}

export { isTemproary, isValid, isExpired, getPayload };