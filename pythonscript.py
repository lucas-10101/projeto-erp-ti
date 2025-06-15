import os

def juntar_arquivos_em_outtxt(diretorio, arquivo_saida='out.txt'):
    arquivos = [f for f in os.listdir(diretorio) if os.path.isfile(os.path.join(diretorio, f))]
    conteudo_total = []

    for nome_arquivo in arquivos:
        caminho_completo = os.path.join(diretorio, nome_arquivo)
        with open(caminho_completo, 'r', encoding='utf-8') as f:
            conteudo = f.read()
            conteudo_total.append(conteudo)
    
    with open(arquivo_saida, 'w', encoding='utf-8') as out:
        out.write(('\n#####\n').join(conteudo_total))

# Exemplo de uso:
# Substitua '.' pelo caminho do diretório desejado
juntar_arquivos_em_outtxt('/home/lucas/repositorios/projeto-erp-ti/organization-api/data/entities')
