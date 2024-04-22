import random for panda


#abyss, legacy


# BACKUP  AP2

#Auto-Avaliação
#Beatriz - TA
#Isabella - TA
#Hanaah - TA
#Luigi - TA
#Luiza - TA
#Joao - NT

# SISTEMA SISLOJA

estoque = []
listacodigo = []
listavalores = []
listaquantidade = []
listaCPF = set()
listaRenda_clientes = []
listaGeral = dict()
clientes = []
caixa = []
modulo = []

# função do GerEstoque
def f(estoque):

    def i(estoque):  # funcao da inclusao do estoque

        quantidade = int(input("Quantos produtos serão incluídos? "))
        soma = 0


        for i in range(quantidade):

            while True:
              cod_item = input("Digite o codigo do item: ")
              for digito in cod_item:
                soma += int(digito)
              if soma >= 30 and soma <= 100:
                desc_item = input("Digite a descrição do item: ")
                valor_item = float(input("Digite o valor em reais do item: "))
                print("")
                estoque.append((cod_item, desc_item, valor_item))
                listavalores.append(valor_item)
                listacodigo.append(cod_item)
                break

              else:
                print("Código invalido")


        media = sum(listavalores) / len(listavalores) #calcular a media dos valores dos itens
        maior_valor = max([float(i) for i in listavalores])
        print('----------------------------------------------------')
        print(f"Valor médio dos itens cadastrados: {media:.2f}")
        print(maior_valor)
        for valor_item, cod_item in zip(listavalores, listacodigo): # associa uma lista a outra, achando o elemento correspondente
            if float(valor_item) == maior_valor:
                print(f"Item de maior valor cadastrado: {cod_item} (valor: {maior_valor:.2f})")
                print('-----------------------------------------------------')



        while True:  # opcao para o usuario retornar a tela principal ou utilizar novamente o modulo estoque
            menu = int(input("\n[0] Para retornar à tela principal \n[1] Para iniciar o estoque novamente."))
            if (menu == 0):
                m(modulo)
            elif (menu == 1):
                f(estoque)


    def e(estoque):  # funcao da exclusao do estoque
        soma = 0

        print("Lista atual: ", estoque)
        codigo = input("Digite os código dos itens que deseja retirar, separados por vírgula: ")
        for digito in codigo: #faz a soma dos digitos
          soma += int(digito)
        while True:
            if soma >= 30 and soma <= 100: #verifica se a soma dos digitos é maior ou igual a 30 ou menor ou igual a 100
              c = input("Voce tem certeza que deseja excluir o item?\n[1] Sim \n[2] Não \n")
              match c:
                  case '1':
                      exclusao = False

                      for item in estoque: # procura o item no estoque
                          if item[0] == codigo:
                              item_removido = item
                              estoque.remove(item)
                              exclusao = True #assim que achar o item retorna o valor True para entrar no if

                              break

                      if exclusao:
                        nova_ex = input("Deseja realizar uma nova exclusão?: ")
                        match nova_ex:
                            case '1':
                              e(estoque)

                            case '2':


                              print("Operação realizada com sucesso")
                              print('------------------')
                              print('relatório dos itens excluídos')
                              print('Código     Saldo')
                              print(item_removido[0], "     ",item_removido[2])

                      else:
                        print("CÓDIGO INVALIDO")


                      while True:  # opcao para o usuario retornar a tela principal ou utilizar novamente o modulo estoque
                          menu = int(input("\n[0] Para retornar à tela principal \n[1] Para iniciar o estoque novamente."))
                          if (menu == 0):
                              m(modulo)
                          elif (menu == 1):
                              f(estoque)
                  case '2':
                      m(modulo)
              break


            else:
              print("Codigo invalido digite novamente")
              e(estoque)
    opcoesestoque = input("[1] Para inclusão \n[2] Para exclusão \n : ")
    match (opcoesestoque):
       case '1':
         i(estoque)
       case '2':
          e(estoque)

#--------------------------------------------------------------------------------------

# GerClientes
def g(clientes):
    # Função para encontrar o primeiro dígito verificador
    def d1(cpf_d1):
        num_str = str(cpf_d1)      # converter para uma string
        primeiros_nove = num_str[0:9]      # selecionar os primeiros nove dígitos
        lista_numeros = [int(digito) for digito in primeiros_nove]      # converter cada dígito em um número inteiro
        lista_multiplicada = [lista_numeros[i] * (10 - i) for i in range(len(lista_numeros))]        # aplicar a multiplicação
        soma = sum(lista_multiplicada)         # somar os resultados
        r1 = soma % 11         # obter o resto da divisão por 11
        if r1 >= 2:
            d1 = 11 - r1
        else:
            d1 = 0
        r1 = str(d1)
        return r1

    # Função para encontrar o segundo dígito verificador
    def d2(cpf_d2):
        num_str = str(cpf_d2)      # converter para uma string
        primeiros_nove = num_str[1:10]      # pular o primeiro e selecionar os outros nove dígitos
        lista_numeros = [int(digito) for digito in primeiros_nove]      # converter cada dígito em um número inteiro
        lista_multiplicada = [lista_numeros[i] * (10 - i) for i in range(len(lista_numeros))]          # aplicar a multiplicação
        soma = sum(lista_multiplicada)        # somar os resultados
        r2 = soma % 11         # obter o resto da divisão por 11
        if r2 >= 2:
            d2 = 11 - r2
        else:
            d2 = 0
        r2 = str(d2)
        return r2

        # Função para encontar o último e o penúltimo do CPF
    def penultimo_e_ultimo(peun):
        peun_inteiro = int(peun)        # converter para um inteiro
        ultimo_digito = peun_inteiro % 10         # obter o resto da divisão por 10
        penultimo_digito = (peun_inteiro // 10) % 10
        str_ultimo = str(ultimo_digito)           # converter para uma string
        str_penultimo = str(penultimo_digito)          # converter para uma string
        penultimo_e_ultimo = str_penultimo + str_ultimo
        return penultimo_e_ultimo

    def clientes(listaRenda_clientes):
        clientes_ate_5 = sum(1 for renda in listaRenda_clientes if renda < 5000) # confere as rendas inferiores a 5000 reais
        clientes_entre_5e10 = sum(1 for renda in listaRenda_clientes if 5000 <= renda < 10000) # confere as rendas inferiores a 10000 reais e inferiores a 5000 reais
        clientes_acima_10 = sum(1 for renda in listaRenda_clientes if renda >= 10000) # confere as rendas superiores a 5000 reais
        total_clientes = clientes_ate_5 + clientes_entre_5e10 + clientes_acima_10
        print('Total de clientes cadastrados: {}'.format(total_clientes))

        print('------------------------------------------------') # interface
        print('FAIXA                             PORCENTAGEM')

        percentual_5 = (clientes_ate_5 / total_clientes) * 100
        print('Até R$5.000                        {:.2f}%'.format(percentual_5))
        percentual_5a10 = (clientes_entre_5e10 / total_clientes) * 100
        print('Entre R$5.000 e R$10.000           {:.2f}%'.format(percentual_5a10))
        percentual_10 = (clientes_acima_10 / total_clientes) * 100
        print('Acima de R$10.000                  {:.2f}%'.format(percentual_10))

        print('------------------------------------------------')
        return (" ")

    print('Para parar a operação aperte o 0 uma vez.\n')

    flag = 1
    while flag == 1:
        cpf = str(input('Digite o seu CPF (Somente números): '))
        cpf_inteiro = float(cpf)
        if cpf_inteiro == 0:
                break

        dois_digitos_verificados = d1(cpf) + d2(cpf)

        if dois_digitos_verificados == penultimo_e_ultimo(cpf):
            print('CPF VALIDO')

            listaCPF.add(cpf) # adiciona o CPF do cliente na lista de CPF

            if cpf in listaCPF:
              renda_cliente = float(input('Digite sua renda (Somente números): '))
              listaGeral[cpf] = renda_cliente
              listaRenda_clientes.append(renda_cliente) # adiciona a renda na lista de rendas

        else:
            print('CPF INVÁLIDO')

        print('\n')

    print('\n')

    print('------------------------------------') # interface
    print('OPERAÇÃO REALIZADA COM SUCESSO')
    print(clientes(listaRenda_clientes)) # mostra a lista de rendas

    while True:
      sim_nao = int(input('Digite 1 caso queira ver a sua renda e 2 caso não queira ver a sua renda: '))
      if sim_nao == 1:
          encontrar_cpf = input('Digite o seu CPF: ')
          if encontrar_cpf in listaGeral:
              print('Renda: ',listaGeral[encontrar_cpf])
              break

          else:
              print('Saindo')
              break

      elif sim_nao == 2:
          print('Saindo')
          break

      else:
        print('Erro')

      print('\n')

    print('\n')

    while True:  # opcao para o usuario retornar a tela principal ou utilizar novamente o modulo clientes
        menu = int(input("[0] Para retornar à tela principal \n[1] Para iniciar o clientes novamente.\n"))
        if (menu == 0):
            m(modulo)
        elif (menu == 1):
            g(clientes)


# ---------------------------------------------------------------------------

# funcao do gercaixa
def c(caixa):
    print(estoque)  # mostra o estoque
    print(listaCPF)

    while True:
        print("[1] Transação de Movimentação Financeira")
        print("[2] Cadastrar Venda")
        print("[3] Quantidade de Vendas em Determinado Dia")
        print("[0] Retornar à tela principal")


        opcao = input("Escolha uma opção: ")

        if opcao == "1":
            def quantidade_vendida_dia(data):
                quant_total = 0
                for venda in listacaixa:
                    if venda['data'] == data:
                        quant_total += venda['quant_item']
                return quant_total

            data = input("\nData [dd/mm/aaaa]: ")
            saldo_inicial = float(input("Saldo inicial do caixa no dia: "))
            N = int(input("Número total de vendas no dia: "))
            listacaixa = []
            for i in range(N):
                flag = True
                while flag:
                    cpf = input("CPF: ")
                    if cpf in listaCPF:
                      flag = False
                    else:
                        print("CPF não está incluso na lista de Clientes.\n")

                flag = True
                while flag:
                    cod_item = int(input("Código do item: "))
                    if cod_item in estoque:
                      flag = False
                    else:
                        print("Código não faz parte do estoque.")
                quant_item = int(input("Quantidade de itens: "))
                valor_un = float(input("Valor do item: "))

                listacaixa.append({'data': data, 'cpf': cpf, 'cod_item': cod_item, 'quant_item': quant_item, 'valor_un': valor_un})

            tot_venda = 0
            for venda in listacaixa:
                tot_venda += venda['quant_item'] * venda['valor_un']

            med_venda = tot_venda / N
            saldo_final = saldo_inicial + tot_venda

            print("\n")
            print("----------------------------------------")
            print("RELATÓRIO DE MOVIMENTAÇÃO FINANCEIRA")
            print("Data da movimentação:", data)
            print("Saldo: R$", saldo_final)
            print("Valor médio das vendas: R$", med_venda)
            print("Total das vendas:", N, "unidades")
            print("----------------------------------------")
            print("\n")

        elif opcao == "2":
            def cadastro_vendas(data, cpf, cod_item, quant_item):
                cad_venda = {
                    'data': data,
                    'cpf': cpf,
                    'cod_item': cod_item,
                    'quant_item': quant_item
                }
                listacaixa.append(cad_venda)

            data = input("Data da venda [dd/mm/aaaa]: ")

            flag = True
            while flag:
                cpf = input("CPF: ")
                if cpf in listaCPF:
                  flag = False
                else:
                  print("CPF não está incluso na lista de Clientes.\n")

            flag = True
            while flag:
                cod_item = int(input("Código do item: "))
                if cod_item in estoque:
                  flag = False
                else:
                  print("Código não faz parte do estoque.")

            quant_item = int(input("Quantidade de itens: "))

            cadastro_vendas(data, cpf, cod_item, quant_item)

            print("\n")
            print("----------------------------------------")
            print("CADASTRO DA VENDA")
            print("Data da movimentação:", data)
            print("Código do cliente:", cpf)
            print("Código do item:", cod_item)
            print("Total das vendas:", quant_item, "unidades")
            print("----------------------------------------")
            print("\n")

        elif opcao == "3":
            data = input("Data [dd/mm/aaaa]: ")
            quantidade = quantidade_vendida_dia(data)
            print("\n")
            print("----------------------------------------")
            print("Quantidade de vendas no dia", data, ":", quantidade, "unidades")
            print("----------------------------------------")
            print("\n")

        elif opcao == "0":
            c(caixa)
            break

        else:
            print("Opção inválida. Por favor, escolha uma opção válida.")

    while True:  # opcao para o usuario retornar a tela principal ou utilizar novamente o modulo caixa
        menu = int(input("[0] Para retornar à tela principal \n[1] Para iniciar o caixa novamente.\n"))
        if (menu == 0):
            m(modulo)
        elif (menu == 1):
            c(caixa)


# ---------------------------------------------------------------------------


# Função menu
def m(modulo):
    print("Bem vindo ao sistema SISLOJA!")
    while True:
        modulo = input(
            "Escolha o módulo desejado:\n [1] Para GerEstoque \n [2] Para GerClientes\n [3] Para GerCaixa \n [4] Para sair \nDigite o módulo que deseja: ")
        print("\n")

        match modulo:

            case '1':
                f(estoque)
            case "2":
                g(clientes)  # erro no cpf
            case "3":
                c(caixa)  # ver se precisa verificar o cpf

            case "4":
                print("Saindo")
                break

    return ("Esse foi o sistema SISLOJA")


m(modulo)
