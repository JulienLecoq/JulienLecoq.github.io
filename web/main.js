function copyToClipboard(text, onSuccess, onError) {
    navigator.clipboard.writeText(text).then(onSuccess).catch(onError);
}
